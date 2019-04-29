package server

import (
	"bytes"
	"encoding/binary"
	"gitee.com/johng/gf/g/os/glog"
	"gitee.com/johng/gf/g/os/grpool"
	"github.com/miekg/dns"
	"godial/client"
	"godial/common"
	"godial/gen-go/rpc/yamutech/com"
	"net"
	"os"
	"sync"
	"time"
)

var Pool *grpool.Pool

const PoolSize = 100000

func NewPool() {
	Pool = grpool.New(PoolSize)
}

type DomainResult struct {
	result []*com.DomainResult_
	lock   *sync.Mutex
}

func NewDomainResult() *DomainResult {
	result := &DomainResult{}
	result.lock = new(sync.Mutex)
	return result
}

func (r *DomainResult) Append(val ...*com.DomainResult_) {
	r.lock.Lock()
	r.result = append(r.result, val...)
	r.lock.Unlock()
}

func (r *DomainResult) Len() int {
	r.lock.Lock()
	length := len(r.result)
	r.lock.Unlock()
	return length
}

func cal_mask(val int32) uint32 {
	var i int32 = 0
	var res uint32 = 1
	for i = 0; i < val; i++ {
		res *= 2
	}
	res -= 1
	return ^res
}

func IntToBytes(n uint32) []byte {
	x := uint32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func Ipv4Local(a net.IP) (bool, bool) {
	var local, res bool
	IpsecMap.Iterator(func(k string, v interface{}) bool {
		ipsec, ok := v.(*com.IpSec)
		if ok {
			mask := IntToBytes(cal_mask(32 - ipsec.Mask))
			ipnet := net.IPNet{net.ParseIP(ipsec.IP.Addr), net.IPv4Mask(mask[0], mask[1], mask[2], mask[3])}
			if ipnet.Contains(a) {
				local = ipsec.Local
				res = true
				return false
			}
		}
		return true
	})
	return local, res
}

func Ipv4PolicyDial(dname string, a *dns.A, method com.DialMethod, ip_result *com.IpResult_) bool {
	var dname_local bool
	ip_result.IP = com.NewIpAddr()
	ip_result.IP.Addr = a.A.String()
	ip_result.IP.Version = 4
	ip_result.Local, dname_local = Ipv4Local(a.A)

	switch method {
	case com.DialMethod_Dig:
		ip_result.Available = true
	case com.DialMethod_DigAndPing:
		ip_result.Delay, ip_result.Available = Ping(a.A)
	case com.DialMethod_DigAndHttp:
		ip_result.Delay, ip_result.Available = HttpDial(a.A)
	case com.DialMethod_DigAndWeb:
		HttpHttpsDial(dname, a.A, ip_result)
	case com.DialMethod_DigAndVideo:
		HttpHttpsDial(dname, a.A, ip_result)
	case com.DialMethod_DomainSchedul:
		HttpHttpsDial(dname, a.A, ip_result)
	}

	return dname_local
}

func Ipv6PolicyDial(dname string, a *dns.AAAA, method com.DialMethod, ip_result *com.IpResult_) bool {
	var dname_local bool
	ip_result.IP = com.NewIpAddr()
	ip_result.IP.Addr = a.AAAA.String()
	ip_result.IP.Version = 6
	//ip_result.Local, dname_local = Ipv6Local(a.A)

	switch method {
	case com.DialMethod_Dig:
		ip_result.Available = true
	case com.DialMethod_DigAndPing:
		ip_result.Delay, ip_result.Available = Ping(a.AAAA)
	case com.DialMethod_DigAndHttp:
		ip_result.Delay, ip_result.Available = HttpDial(a.AAAA)
	case com.DialMethod_DigAndWeb:
		WebDial(dname, a.AAAA, ip_result)
	case com.DialMethod_DigAndVideo:
		WebDial(dname, a.AAAA, ip_result)
	case com.DialMethod_DomainSchedul:
		HttpHttpsDial(dname, a.AAAA, ip_result)
	}

	return dname_local
}

func DigAndDial(dname, server string, domaintype uint16, method com.DialMethod, dname_result *com.DomainResult_) {
	dname_result.Dname = dname
	dname_result.Fdr = com.NewFocusDomainResult_()
	dname_result.Dtype = com.DomainType(domaintype)
	server = server + ":53"
	remsg, n, dns_delay := SendDnsPkg(server, dname, domaintype)
	if n > 0 {
		msg := new(dns.Msg)
		err := msg.Unpack(remsg)
		if err == nil {
			dname_result.Available = true
			dname_result.Delay = int32(dns_delay.Nanoseconds()) / 1000
			dname_result.Fdr.Delay = dname_result.Delay
			if rcode := msg.MsgHdr.Rcode; rcode > 5 || rcode < 0 {
				dname_result.Fdr.Status = com.FocusDomainResultStatus_others
			} else {
				dname_result.Fdr.Status = com.FocusDomainResultStatus(rcode)
			}

			for i := 0; i < len(msg.Answer); i++ {
				switch msg.Answer[i].Header().Rrtype {
				case dns.TypeA:
					rr := msg.Answer[i]
					a, ok := rr.(*dns.A)
					if ok {
						// result 赋值 拨测ping/http/web等
						ip_result := com.NewIpResult_()
						dname_result.Local = Ipv4PolicyDial(dname, a, method, ip_result)
						dname_result.Results = append(dname_result.Results, ip_result)
						item_result := com.NewFocusDomainResultItem()
						item_result.Value = a.A.String()
						dname_result.Fdr.Results = append(dname_result.Fdr.Results, item_result)
					}
				case dns.TypeAAAA:
					rr := msg.Answer[i]
					aaaa, ok := rr.(*dns.AAAA)
					if ok {
						ip_result := com.NewIpResult_()
						dname_result.Local = Ipv6PolicyDial(dname, aaaa, method, ip_result)
						dname_result.Results = append(dname_result.Results, ip_result)
						item_result := com.NewFocusDomainResultItem()
						item_result.Value = aaaa.AAAA.String()
						dname_result.Fdr.Results = append(dname_result.Fdr.Results, item_result)
					}
				case dns.TypeMX:
					rr := msg.Answer[i]
					mx, ok := rr.(*dns.MX)
					if ok {
						item_result := com.NewFocusDomainResultItem()
						item_result.Priority = int32(mx.Preference)
						item_result.Value = mx.Mx
						dname_result.Fdr.Results = append(dname_result.Fdr.Results, item_result)
					}
				case dns.TypeNS:
					rr := msg.Answer[i]
					ns, ok := rr.(*dns.NS)
					if ok {
						item_result := com.NewFocusDomainResultItem()
						item_result.Value = ns.Ns
						dname_result.Fdr.Results = append(dname_result.Fdr.Results, item_result)
					}
				case dns.TypeCNAME:
					rr := msg.Answer[i]
					cname, ok := rr.(*dns.CNAME)
					if ok {
						item_result := com.NewFocusDomainResultItem()
						item_result.Value = cname.Target
						dname_result.Fdr.Results = append(dname_result.Fdr.Results, item_result)
					}
				case dns.TypePTR:
					rr := msg.Answer[i]
					ptr, ok := rr.(*dns.PTR)
					if ok {
						item_result := com.NewFocusDomainResultItem()
						item_result.Value = ptr.Ptr
						dname_result.Fdr.Results = append(dname_result.Fdr.Results, item_result)
					}
				}
			}
		}
	}
}

func DoTask(task *Task, domain com.DomainRecord, result *DomainResult) {
	domain_result := com.NewDomainResult_()
	DigAndDial(domain.Dname, task.target_table.Addr, uint16(domain.Dtype), task.method, domain_result)
	result.Append(domain_result)
	//glog.Println("task :", domain_result)
}

func report_task_idle(id string) {
	task_process := com.NewTaskProcessArgs_()
	task_process.Event = com.TaskEvent_IDLE
	task_process.Batchno = time.Now().Format("2006-01-02 15:04:05")
	if TaskMap.Contains(id) {
		client.Client.Lock.Lock()
		client.Client.ReportClient.ReportTaskProcess(common.ModuleId, id, task_process)
		client.Client.Lock.Unlock()
		glog.Println("task idle: taskid =", id)
	}
}

func report_task_plan(task_process *com.TaskProcessArgs_, id string, interval int32, rcv_cnt, dname_cnt int) {
	task_process.Event = com.TaskEvent_RUNNING
	task_process.Closed = (0 == interval)
	task_process.Batchno = time.Now().Format("2006-01-02 15:04:05")
	task_process.Percent = (float64(rcv_cnt) * 100 / float64(dname_cnt))
	if TaskMap.Contains(id) {
		client.Client.Lock.Lock()
		client.Client.ReportClient.ReportTaskProcess(common.ModuleId, id, task_process)
		client.Client.Lock.Unlock()
		glog.Println("task runing: taskid =", id, task_process)
	}
}

func report_result(dname_result []*com.DomainResult_, task_process *com.TaskProcessArgs_, id string) {
	reslen := len(dname_result)
	var rcv_cnt, total_dig_delay, total_dial_delay, detect_available, local_dial, local_detect int32
	var available, local bool

	for i := 0; i < reslen; i++ {
		if dname_result[i].Available {
			rcv_cnt++
			total_dig_delay += dname_result[i].Delay
			local, available = true, true
			if iplen := len(dname_result[i].Results); iplen > 0 {
				var dial_delay int32
				for j := 0; j < iplen; j++ {
					dial_delay += dname_result[i].Results[j].Delay
					if dname_result[i].Results[j].Available == false {
						available = false
						dname_result[i].Available = false
					}
					if dname_result[i].Results[j].Local == false {
						local = false
					}
				}
				total_dial_delay += (dial_delay / int32(iplen))
			}
			if available {
				detect_available++
			}
			if local {
				local_detect++
			}
			if dname_result[i].Local {
				local_dial++
			}
		}
		//glog.Println(dname_result[i])
	}

	task_process.Event = com.TaskEvent_FINISHED
	task_process.Batchno = time.Now().Format("2006-01-02 15:04:05")
	task_process.Percent = 100
	if rcv_cnt > 0 {
		task_process.DialLocalRate = float64(local_dial) * 100 / float64(rcv_cnt)
		task_process.DetectLocalRate = float64(local_detect) * 100 / float64(rcv_cnt)
		task_process.DetectAvailRate = float64(detect_available) * 100 / float64(reslen)
		task_process.DialAvgDelay = total_dig_delay / rcv_cnt
		task_process.DetectAvgDelay = total_dial_delay / rcv_cnt
		task_process.TotalAvgDelay = task_process.DialAvgDelay + task_process.DetectAvgDelay
	}

	if TaskMap.Contains(id) {
		client.Client.Lock.Lock()
		cnt := reslen / 2000
		var i int
		if cnt > 0 {
			for i = 0; i < cnt; i++ {
				client.Client.ReportClient.ReportResult_(common.ModuleId, id, task_process.Batchno, dname_result[2000*i:2000*(i+1)])
			}
			if reslen%2000 != 0 {
				client.Client.ReportClient.ReportResult_(common.ModuleId, id, task_process.Batchno, dname_result[2000*i:])
			}
		} else {
			client.Client.ReportClient.ReportResult_(common.ModuleId, id, task_process.Batchno, dname_result)
		}
		client.Client.ReportClient.ReportTaskProcess(common.ModuleId, id, task_process)
		client.Client.Lock.Unlock()
		glog.Println("task finished: taskid =", id, task_process)
	}
}

func RunTask(id string, task *Task) {
	task_process := com.NewTaskProcessArgs_()
	result := NewDomainResult()
	tmp := DomainGroupMap.Get(task.group_id)
	group, ok := tmp.(DomainGroup)
	if ok {
		//上报 任务进度
		report_task_plan(task_process, id, task.interval, result.Len(), group.dname_map.Size())

		group.dname_map.Iterator(func(k string, v interface{}) bool {
			domain, ok := v.(*com.DomainRecord)
			if ok {
				//glog.Println("task :", common.MaxGoCount)
				for {
					if common.GoCount.Add(1) < common.MaxGoCount {
						go DoTask(task, *domain, result)
						break
					} else {
						common.GoCount.Add(-1)
						time.Sleep(time.Second)
					}
				}
			}
			return true
		})

		var wait_time int
		for {
			//glog.Println(result.Len(), group.dname_map.Size())
			if result.Len() < group.dname_map.Size() {
				time.Sleep(time.Second)
				wait_time++
				if wait_time == 5 {
					report_task_plan(task_process, id, task.interval, result.Len(), group.dname_map.Size())
					wait_time = 0
				}
			} else {
				common.GoCount.Add(-group.dname_map.Size())
				//统计dname_result上报 域名结果 任务结果
				report_result(result.result, task_process, id)
				break
			}
		}
	}
	task.run = false
}

func Monitor() {
	var err error
	for {
		TaskMap.Iterator(func(k string, v interface{}) bool {
			t_now := time.Now().Unix()
			task, ok := v.(*Task)
			if ok && task.interval > 0 {
				if ((t_now - task.t_start) >= 0) && task.run == false {
					task.t_start = t_now + int64(task.interval)
					task.run = true
					go RunTask(k, task)
				} else {
					report_task_idle(k)
				}
			}
			return true
		})
		time.Sleep(5 * time.Second)
		client.Client.Lock.Lock()
		_, err = client.Client.RegisterClient.HeartBeat(common.ModuleId)
		client.Client.Lock.Unlock()
		if err != nil {
			glog.Println("HeartBeat error", err)
			os.Exit(1)
		}
	}
}
