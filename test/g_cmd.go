package main

import (
	"fmt"
	"gitee.com/johng/gf/g/os/gcmd"
	"time"
)

func help() {
	fmt.Println("This is help.")
}

func test() {
	fmt.Println("This is test.")
}

func hello() {
	for {
		fmt.Println("This is hello.")
		time.Sleep(time.Second)
	}
}

func main() {
	gcmd.BindHandle("help", help)
	gcmd.BindHandle("test", test)
	go hello()
	gcmd.AutoRun()
}
