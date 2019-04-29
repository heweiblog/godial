package server

import (
	"errors"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gitee.com/johng/gf/g/container/gmap"
	"gitee.com/johng/gf/g/os/glog"
	"godial/common"
	"godial/gen-go/rpc/yamutech/com"
	"os"
	"time"
)

type Handle struct {
	com.Dial
}

type DomainGroup struct {
	dname_map *gmap.StringInterfaceMap
}

type Task struct {
	method       com.DialMethod
	target_table *com.IpAddr
	source_ip    *com.IpAddr
	interval     int32
	t_start      int64
	group_id     string
	run          bool
}

var IpsecMap *gmap.StringInterfaceMap
var DomainGroupMap *gmap.StringInterfaceMap
var TaskMap *gmap.StringInterfaceMap

func NewTask() *Task {
	return &Task{}
}

func (h Handle) HeartBeat() (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) ResetModule() (r com.RetCode, err error) {
	return com.RetCode_OK, nil
}

func (h Handle) AddIpSec(ipSecList []*com.IpSec) (r com.RetCode, err error) {
	for i := 0; i < len(ipSecList); i++ {
		IpsecMap.Set(ipSecList[i].IP.Addr, ipSecList[i])
		glog.Println("add ipsec:", ipSecList[i].IP.Addr, ipSecList[i].Mask, ipSecList[i].Local)
	}
	return com.RetCode_OK, nil
}

func (h Handle) RemoveIpSec(ipSecList []*com.IpSec) (r com.RetCode, err error) {
	for i := 0; i < len(ipSecList); i++ {
		IpsecMap.Remove(ipSecList[i].IP.Addr)
		glog.Println("del ipsec:", ipSecList[i].IP.Addr, ipSecList[i].Mask, ipSecList[i].Local)
	}
	return com.RetCode_OK, nil
}

func (h Handle) ClearIpSec() (r com.RetCode, err error) {
	IpsecMap.Clear()
	glog.Println("clear ipsec")
	return com.RetCode_OK, nil
}

func (h Handle) AddDialDomain(groupId string, DomainList []*com.DomainRecord) (r com.RetCode, err error) {
	var group DomainGroup

	res := DomainGroupMap.Contains(groupId)
	if !res {
		group.dname_map = gmap.NewStringInterfaceMap()
	} else {
		var ok bool
		gp := DomainGroupMap.Get(groupId)
		group, ok = gp.(DomainGroup)
		if !ok {
			glog.Println("add domain_group failed: interface to struct error", groupId)
			return com.RetCode_FAIL, errors.New("interface to struct error")
		}
	}
	for i := 0; i < len(DomainList); i++ {
		group.dname_map.Set(DomainList[i].Dname, DomainList[i])
		glog.Println("add dname:", DomainList[i].Dname, "->", groupId, "dname_map_len=", group.dname_map.Size())
	}
	if !res {
		DomainGroupMap.Set(groupId, group)
		glog.Println("add a new domain_group:", groupId, "GroupMap_Size=", DomainGroupMap.Size())
	}
	return com.RetCode_OK, nil
}

func (h Handle) RemoveDialDomain(groupId string, DomainList []*com.DomainRecord) (r com.RetCode, err error) {
	glog.Println("del domain start:", groupId, DomainList)
	tmp := DomainGroupMap.Get(groupId)
	group, ok := tmp.(DomainGroup)
	if ok {
		for i := 0; i < len(DomainList); i++ {
			group.dname_map.Remove(DomainList[i].Dname)
			glog.Println("del dname:", DomainList[i].Dname, "->", groupId, "dname_map_len=", group.dname_map.Size())
		}
		return com.RetCode_OK, nil
	} else {
		glog.Println("del domain_group failed:", groupId)
		return com.RetCode_FAIL, nil
	}
}

func (h Handle) ClearDialDomain(groupId string) (r com.RetCode, err error) {
	glog.Println("del and clear domain_group start:", groupId)
	tmp := DomainGroupMap.Get(groupId)
	group, ok := tmp.(DomainGroup)
	if ok {
		group.dname_map.Clear()
		DomainGroupMap.Remove(groupId)
		glog.Println("del and clear domain_group success:", groupId, "GroupMap_Size=", DomainGroupMap.Size())
		return com.RetCode_OK, nil
	} else {
		glog.Println("del and clear domain_group failed:", groupId)
		return com.RetCode_FAIL, errors.New("not find dname group")
	}
}

func (h Handle) AddDialTask(taskId string, method com.DialMethod, targetList []*com.IpAddr, sourceip *com.IpAddr, interval int32, domainGroupId string) (r com.RetCode, err error) {

	if len(targetList) == 0 {
		glog.Println("add task failed: task id =", taskId, targetList)
		return com.RetCode_FAIL, errors.New("dst_ip is nil")
	}

	task := NewTask()
	task.method = method
	task.target_table = targetList[0]
	task.source_ip = sourceip
	task.interval = interval
	task.t_start = time.Now().Unix()
	task.group_id = domainGroupId
	TaskMap.Set(taskId, task)

	glog.Println("add task success:", *task)

	return com.RetCode_OK, nil
}

func (h Handle) RemoveDialTask(taskId string) (r com.RetCode, err error) {
	ok := TaskMap.Contains(taskId)
	if ok {
		TaskMap.Remove(taskId)
		glog.Println("del task success", taskId)
		return com.RetCode_OK, nil
	} else {
		glog.Println("del task failed", taskId)
		return com.RetCode_FAIL, nil
	}
}

func Dial_server_start() {
	serverTransport, err := thrift.NewTServerSocket(common.DialAddr)
	if err != nil {
		glog.Println("Error!", err)
		os.Exit(1)
	}
	processor := com.NewDialProcessor(Handle{})
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	glog.Println("Dial thrift server in", common.DialAddr)

	server.Serve()
}

func init() {
	IpsecMap = gmap.NewStringInterfaceMap()
	DomainGroupMap = gmap.NewStringInterfaceMap()
	TaskMap = gmap.NewStringInterfaceMap()
}
