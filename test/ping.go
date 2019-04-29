package main

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}

func ping(ip string) {
	var (
		icmp  ICMP
		laddr = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		//raddr, _ = net.ResolveIPAddr("ip", "192.168.6.55")
		//raddr, _ = net.ResolveIPAddr("ip", "1.1.8.8")
		//raddr, _ = net.ResolveIPAddr("ip", "1.1.8.111")
		raddr, _ = net.ResolveIPAddr("ip", ip)
	)

	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()

	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)

	fmt.Printf("\n正在 Ping %s 具有 0 字节的数据:\n", raddr.String())
	recv := make([]byte, 1024)

	statistic := list.New()
	sended_packets := 0

	for i := 4; i > 0; i-- {

		if _, err := conn.Write(buffer.Bytes()); err != nil {
			fmt.Println(err.Error())
			return
		}
		sended_packets++
		t_start := time.Now()

		conn.SetReadDeadline((time.Now().Add(time.Second * 2)))
		recv_len, err := conn.Read(recv)

		if err != nil {
			fmt.Println("请求超时")
			continue
		}

		t_end := time.Now()

		dur := t_end.Sub(t_start).Nanoseconds() / 1e6

		fmt.Println("recv len =", recv_len)
		fmt.Println("recv =\n", recv)
		fmt.Printf("来自 %s 的回复: 时间 = %dms\n", raddr.String(), dur)

		statistic.PushBack(dur)

		//for i := 0; i < recvsize; i++ {
		//	if i%16 == 0 {
		//		fmt.Println("")
		//	}
		//	fmt.Printf("%.2x ", recv[i])
		//}
		//fmt.Println("")

	}
	defer func() {
		fmt.Println("")
		//信息统计
		var min, max, sum int64
		if statistic.Len() == 0 {
			min, max, sum = 0, 0, 0
		} else {
			min, max, sum = statistic.Front().Value.(int64), statistic.Front().Value.(int64), int64(0)
		}

		for v := statistic.Front(); v != nil; v = v.Next() {

			val := v.Value.(int64)

			switch {
			case val < min:
				min = val
			case val > max:
				max = val
			}

			sum = sum + val
		}
		recved, losted := statistic.Len(), sended_packets-statistic.Len()
		fmt.Printf("%s 的 Ping 统计信息：\n  数据包：已发送 = %d，已接收 = %d，丢失 = %d (%.1f%% 丢失)，\n往返行程的估计时间(以毫秒为单位)：\n  最短 = %dms，最长 = %dms，平均 = %.0fms\n",
			raddr.String(),
			sended_packets, recved, losted, float32(losted)/float32(sended_packets)*100,
			min, max, float32(sum)/float32(recved),
		)
	}()
}

func main() {
	//go ping("1.1.1.1")
	//go ping("1.1.8.8")
	//go ping("8.8.8.8")
	//go ping("149.20.1.49")
	go ping("192.168.6.195")
	//go ping("1.1.2.3")
	time.Sleep(time.Second * 10)
}
