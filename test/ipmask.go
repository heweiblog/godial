package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.168.6.22")
	fmt.Println(ip)
	ipmask := net.IPv4Mask(255, 255, 255, 0)
	fmt.Println(ipmask)
	ipnet := net.IPNet{ip, ipmask}
	fmt.Println(ipnet)

	ip1 := net.ParseIP("192.168.6.254")
	fmt.Println(ip1, ipnet, ipnet.Contains(ip1))

	ip2 := net.ParseIP("192.168.6.1")
	fmt.Println(ip2, ipnet, ipnet.Contains(ip2))
}
