package common

import (
	"gitee.com/johng/gf/g/container/gtype"
	//"gitee.com/johng/gf/g/os/glog"
	//"fmt"
	//"github.com/Unknwon/goconfig"
	//"os"
)

const (
	ConfigFile = "dial.ini"
)

var (
	//Config     *goconfig.ConfigFile
	GoCount    *gtype.Int
	ModuleId   int32  = 1
	MaxGoCount int    = 500000
	DialAddr   string = "127.0.0.1:9092"
	AgentIp    string = "127.0.0.1"
	DialIp     string = "127.0.0.1"
	AgentPort  int32  = 9196
	DialPort   int32  = 9092
)

/*
func GetConfig() {
	cfg, err := goconfig.LoadConfigFile(ConfigFile)
	if err != nil {
		os.Exit(1)
	}
	dialport, err := cfg.GetValue("dial", "port")
	if err != nil {
		os.Exit(1)
	}
	DialIp, err = cfg.GetValue("dial", "ip")
	if err != nil {
		os.Exit(1)
	}
	DialAddr = DialIp + ":" + dialport
	DialPort, err := cfg.Int("dial", "port")
	if err != nil {
		os.Exit(1)
	}

	AgentIp, err = cfg.GetValue("agent", "ip")
	if err != nil {
		os.Exit(1)
	}
	AgentPort, err := cfg.Int("agent", "port")
	if err != nil {
		os.Exit(1)
	}
	ModuleId, err := cfg.Int("module", "id")
	if err != nil {
		os.Exit(1)
	}
	MaxGoCount, err := cfg.Int("module", "runmax")
	if err != nil {
		os.Exit(1)
	}

	glog.Println("DialIp:", DialIp, "DialPort:", DialPort)
	glog.Println("AgentIp:", AgentIp, "AgentPort:", AgentPort)
	glog.Println("ModuleId:", ModuleId, "MaxGoCount:", MaxGoCount)
}
*/

func init() {
	GoCount = gtype.NewInt()
}
