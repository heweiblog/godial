package main

import (
	"fmt"
	"github.com/miekg/dns"
	"godial/server"
)

func main() {
	m, n, t := server.SendDnsPkg("1.1.1.1:53", "www.baidu.com", 1)
	//fmt.Println(string(m))
	fmt.Println("size = ", n)
	fmt.Println("delay = ", int32(t.Nanoseconds())/1000)
	msg := new(dns.Msg)
	err := msg.Unpack(m)
	if err == nil {
		fmt.Println(msg)
	}

}
