package main

import (
	"dial/gen-go/rpc/yamutech/com"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
)

func main() {

	ip := "127.0.0.1"
	port := "9999"
	input := "heweiwei very strong"

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	tSocket, err := thrift.NewTSocket(net.JoinHostPort(ip, port))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error resolving address, ", err)
		os.Exit(1)
	}

	tTransport := transportFactory.GetTransport(tSocket)

	client := com.NewDialClientFactory(tTransport, protocolFactory)
	if err := tTransport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, (fmt.Errorf("Error opening socket to %s:%s : %v", ip, port, err)))
		os.Exit(1)
	}
	defer tTransport.Close()

	resp, _ := client.RemoveDialTask(input)
	fmt.Println(resp)
}
