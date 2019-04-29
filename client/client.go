package client

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gitee.com/johng/gf/g/os/glog"
	"godial/common"
	"godial/gen-go/rpc/yamutech/com"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type DialClient struct {
	Lock *sync.Mutex
	//RegisterClient *com.AgentClient
	//ReportClient   *com.AnalysisClient
	RegisterClient  *com.CollectClient
	ReportClient    *com.CollectClient
	ProtocolFactory *thrift.TBinaryProtocolFactory
	Transport       thrift.TTransport
}

var Client *DialClient

func NewClient() {
	var err error
	Client = &DialClient{}
	Client.Transport, err = thrift.NewTSocket(net.JoinHostPort(common.AgentIp, strconv.Itoa(int(common.AgentPort))))
	if err != nil {
		glog.Println(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	Client.ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	Client.Lock = new(sync.Mutex)
	//Client.RegisterClient = com.NewAgentClientFactory(Client.Transport, Client.ProtocolFactory)
	//Client.ReportClient = com.NewAnalysisClientFactory(Client.Transport, Client.ProtocolFactory)
	Client.RegisterClient = com.NewCollectClientFactory(Client.Transport, Client.ProtocolFactory)
	Client.ReportClient = com.NewCollectClientFactory(Client.Transport, Client.ProtocolFactory)
}

func Register() {
	NewClient()

	glog.Println("DialIp:", common.DialIp, "DialPort:", common.DialPort)
	glog.Println("AgentIp:", common.AgentIp, "AgentPort:", common.AgentPort)
	glog.Println("ModuleId:", common.ModuleId, "MaxGoCount:", common.MaxGoCount)

	if err := Client.Transport.Open(); err != nil {
		glog.Println(os.Stderr, "Error opening socket", err)
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)

	ip := com.NewIpAddr()
	ip.Version = 4
	ip.Addr = common.AgentIp
	Client.Lock.Lock()
	for {
		ret, err := Client.RegisterClient.RegisterModule(common.ModuleId, ip, common.DialPort)
		if err != nil {
			glog.Println("register err:", ret, fmt.Errorf("%v", err))
			time.Sleep(time.Second)
			continue
		}
		break
	}
	Client.Lock.Unlock()

	glog.Println("success register ->", common.AgentIp, common.AgentPort)
}
