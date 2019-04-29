package main

import (
	"fmt"
	"net"
	"time"
)

func getpkg(dname, path string) string {
	get := fmt.Sprintf("GET %s HTTP/1.1\r\n", path)
	host := fmt.Sprintf("Host: %s\r\n", dname)
	user_agent := "User-Agent: Mozilla/5.0 (Windows NT 6.2; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0\r\n"
	accept := "Accept:*/*\r\n"
	lan := "Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3\r\n"
	encode := "Accept-Encoding: gzip, deflate\r\n"
	request := get + host + user_agent + accept + lan + encode + "\r\n"
	return request
}

func main() {
	//ip := "61.129.7.47"
	dname := "www.baidu.com"
	path := "/"
	request := getpkg(dname, path)
	conn, err := net.Dial("tcp", "115.239.211.112:80")
	if err != nil {
		fmt.Println("connect error")
		return
	}

	fmt.Println(request, conn.RemoteAddr())
	fmt.Println(conn)
	conn.SetReadDeadline((time.Now().Add(time.Second * 2)))

	n, err := conn.Write([]byte(request))
	if err != nil {
		fmt.Println("connect error")
		return
	}
	fmt.Println(n)

	var rcv []byte
	n, err = conn.Read(rcv)
	if err != nil {
		fmt.Println("connect error")
		return
	}

	fmt.Println(n)

	fmt.Println(string(rcv))
}
