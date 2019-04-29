// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"rpc/yamutech/com"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  RetCode registerModule(i32 moduleId, IpAddr ip, i32 port)")
	fmt.Fprintln(os.Stderr, "  RetCode unRegisterModule(i32 moduleId)")
	fmt.Fprintln(os.Stderr, "  RetCode heartBeat(i32 moduleId)")
	fmt.Fprintln(os.Stderr, "  RetCode reportTaskProcess(i32 moduleId, string taskId, TaskProcessArgs arg)")
	fmt.Fprintln(os.Stderr, "  RetCode reportResult(i32 moduleId, string taskId, string batchno,  resultList)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := com.NewCollectClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "registerModule":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RegisterModule requires 3 args")
			flag.Usage()
		}
		tmp0, err102 := (strconv.Atoi(flag.Arg(1)))
		if err102 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		arg103 := flag.Arg(2)
		mbTrans104 := thrift.NewTMemoryBufferLen(len(arg103))
		defer mbTrans104.Close()
		_, err105 := mbTrans104.WriteString(arg103)
		if err105 != nil {
			Usage()
			return
		}
		factory106 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt107 := factory106.GetProtocol(mbTrans104)
		argvalue1 := com.NewIpAddr()
		err108 := argvalue1.Read(jsProt107)
		if err108 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		tmp2, err109 := (strconv.Atoi(flag.Arg(3)))
		if err109 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.RegisterModule(value0, value1, value2))
		fmt.Print("\n")
		break
	case "unRegisterModule":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UnRegisterModule requires 1 args")
			flag.Usage()
		}
		tmp0, err110 := (strconv.Atoi(flag.Arg(1)))
		if err110 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.UnRegisterModule(value0))
		fmt.Print("\n")
		break
	case "heartBeat":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "HeartBeat requires 1 args")
			flag.Usage()
		}
		tmp0, err111 := (strconv.Atoi(flag.Arg(1)))
		if err111 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.HeartBeat(value0))
		fmt.Print("\n")
		break
	case "reportTaskProcess":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "ReportTaskProcess requires 3 args")
			flag.Usage()
		}
		tmp0, err112 := (strconv.Atoi(flag.Arg(1)))
		if err112 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg114 := flag.Arg(3)
		mbTrans115 := thrift.NewTMemoryBufferLen(len(arg114))
		defer mbTrans115.Close()
		_, err116 := mbTrans115.WriteString(arg114)
		if err116 != nil {
			Usage()
			return
		}
		factory117 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt118 := factory117.GetProtocol(mbTrans115)
		argvalue2 := com.NewTaskProcessArgs_()
		err119 := argvalue2.Read(jsProt118)
		if err119 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.ReportTaskProcess(value0, value1, value2))
		fmt.Print("\n")
		break
	case "reportResult":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "ReportResult_ requires 4 args")
			flag.Usage()
		}
		tmp0, err120 := (strconv.Atoi(flag.Arg(1)))
		if err120 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		arg123 := flag.Arg(4)
		mbTrans124 := thrift.NewTMemoryBufferLen(len(arg123))
		defer mbTrans124.Close()
		_, err125 := mbTrans124.WriteString(arg123)
		if err125 != nil {
			Usage()
			return
		}
		factory126 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt127 := factory126.GetProtocol(mbTrans124)
		containerStruct3 := com.NewCollectReportResultArgs()
		err128 := containerStruct3.ReadField4(jsProt127)
		if err128 != nil {
			Usage()
			return
		}
		argvalue3 := containerStruct3.ResultList
		value3 := argvalue3
		fmt.Print(client.ReportResult_(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
