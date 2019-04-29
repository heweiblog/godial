package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"reflect"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("dial.ini")
	if err != nil {
		panic("错误")
	}
	value, err := cfg.GetValue("dial", "ip")
	if err == nil {
		fmt.Println(reflect.TypeOf(value))
		fmt.Println(value)
	}
	v, err := cfg.Int("agent", "port")
	if err == nil {
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(v)
	}

}
