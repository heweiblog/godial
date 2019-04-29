package main

import (
	"fmt"
	"gitee.com/johng/gf/g/container/gtype"
	"github.com/Unknwon/goconfig"
	"os"
	"reflect"
)

const (
	ConfigFile = "dial.ini"
)

var (
	ModuleId   int
	MaxGoCount int
	DialAddr   string
	AgentIp    string
	DialIp     string
	AgentPort  int
	DialPort   int
)

func GetConfig() {
	cfg, err := goconfig.LoadConfigFile(ConfigFile)
	if err != nil {
		os.Exit(1)
	}
	dialport, err := cfg.GetValue("dial", "port")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(dialport, reflect.TypeOf(dialport))
	DialIp, err = cfg.GetValue("dial", "ip")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(DialIp, reflect.TypeOf(DialIp))
	DialAddr = DialIp + ":" + dialport
	fmt.Println(DialAddr, reflect.TypeOf(DialAddr))
	DialPort, err := cfg.Int("dial", "port")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(DialPort, reflect.TypeOf(DialPort))

	AgentIp, err = cfg.GetValue("agent", "ip")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(AgentIp, reflect.TypeOf(AgentIp))
	AgentPort, err := cfg.Int("agent", "port")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(AgentPort, reflect.TypeOf(AgentPort))
	ModuleId, err := cfg.Int("module", "id")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(ModuleId, reflect.TypeOf(ModuleId))
	MaxGoCount, err := cfg.Int("module", "runmax")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(MaxGoCount, reflect.TypeOf(MaxGoCount))

	fmt.Println(cfg)
}

var GoCount *gtype.Int

func main() {
	GetConfig()
	GoCount = gtype.NewInt()
}
