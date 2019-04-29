package server

import (
	"bytes"
	"container/list"
	"encoding/binary"
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

func Ping(a net.IP) (int32, bool) {
	var (
		icmp     ICMP
		laddr    = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		raddr, _ = net.ResolveIPAddr("ip", a.String())
	)

	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)

	if err != nil {
		return 0, false
	}

	defer conn.Close()

	icmp.Type = 8

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)

	recv := make([]byte, 1024)

	statistic := list.New()
	sended_packets := 0

	for i := 3; i > 0; i-- {

		if _, err := conn.Write(buffer.Bytes()); err != nil {
			return 0, false
		}
		sended_packets++
		t_start := time.Now()

		conn.SetReadDeadline((time.Now().Add(time.Second * 2)))
		_, err := conn.Read(recv)

		if err != nil {
			continue
		}

		t_end := time.Now()

		dur := t_end.Sub(t_start).Nanoseconds() / 1000

		statistic.PushBack(dur)

	}

	if recved := int64(statistic.Len()); recved > 0 {
		var sum int64
		for v := statistic.Front(); v != nil; v = v.Next() {
			sum = sum + v.Value.(int64)
		}
		return int32(sum / recved), true
	}

	return 0, false
}
