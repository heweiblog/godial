package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strings"
	"time"
)

type dnsHeader struct {
	Id                                 uint16
	Bits                               uint16
	Qdcount, Ancount, Nscount, Arcount uint16
}

func (header *dnsHeader) SetFlag(QR uint16, OperationCode uint16, AuthoritativeAnswer uint16, Truncation uint16, RecursionDesired uint16, RecursionAvailable uint16, ResponseCode uint16) {
	header.Bits = QR<<15 + OperationCode<<11 + AuthoritativeAnswer<<10 + Truncation<<9 + RecursionDesired<<8 + RecursionAvailable<<7 + ResponseCode
}

type dnsQuery struct {
	QuestionType  uint16
	QuestionClass uint16
}

func ParseDomainName(domain string) []byte {
	var (
		buffer   bytes.Buffer
		segments []string = strings.Split(domain, ".")
	)
	for _, seg := range segments {
		binary.Write(&buffer, binary.BigEndian, byte(len(seg)))
		binary.Write(&buffer, binary.BigEndian, []byte(seg))
	}
	binary.Write(&buffer, binary.BigEndian, byte(0x00))

	return buffer.Bytes()
}
func Send(dnsServer, domain string) ([]byte, int, time.Duration) {
	requestHeader := dnsHeader{
		Id:      0x0010,
		Qdcount: 1,
		Ancount: 0,
		Nscount: 0,
		Arcount: 0,
	}
	requestHeader.SetFlag(0, 0, 0, 0, 1, 0, 0)

	requestQuery := dnsQuery{
		QuestionType:  1,
		QuestionClass: 1,
	}

	var (
		conn   net.Conn
		err    error
		buffer bytes.Buffer
	)

	if conn, err = net.Dial("udp", dnsServer); err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0), 0, 0
	}
	defer conn.Close()

	binary.Write(&buffer, binary.BigEndian, requestHeader)
	binary.Write(&buffer, binary.BigEndian, ParseDomainName(domain))
	binary.Write(&buffer, binary.BigEndian, requestQuery)

	buf := make([]byte, 1024)
	t1 := time.Now()
	conn.SetDeadline(time.Unix(t1.Unix()+2, 0))
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0), 0, 0
	}
	fmt.Println("debug1")
	length, err := conn.Read(buf)
	fmt.Println("debug2")
	t := time.Now().Sub(t1)
	return buf, length, t
}

func main() {
	//remsg, n, t := Send("114.114.114.114:53", "www.baidu.com")
	remsg, n, t := Send("1.1.8.8:53", "www.baidu.com")
	//remsg, n, t := Send("192.168.6.195:53", "www.baidu.com")
	//remsg, n, t := Send("1.1.8.8:53", "www.dsfiiiiiiiiiiiiiiiiiiiiiiiii.com")
	if n > 0 {
		fmt.Println(remsg, n, t)
		msg := new(dns.Msg)
		ok := msg.Unpack(remsg)
		if ok == nil {
			fmt.Println(msg)
			for i := 0; i < len(msg.Answer); i++ {
				if dns.TypeA == msg.Answer[i].Header().Rrtype {
					rr := msg.Answer[i]
					a, ok := rr.(*dns.A)
					if ok {
						fmt.Println(a.A)
					}
					fmt.Printf("rr = %v\n", rr)
					fmt.Println("Answer = ", msg.Answer[i].String())
					fmt.Println("dns_header = ", msg.Answer[i].Header().String())
					fmt.Printf("rcode = %d\n", msg.MsgHdr.Rcode)
				}
			}
		} else {
			fmt.Println("pack error")
		}
	}
}
