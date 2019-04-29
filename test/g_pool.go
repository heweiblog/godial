package main

import (
	"fmt"
	"gitee.com/johng/gf/g/os/grpool"
	"gitee.com/johng/gf/g/os/gtime"
	"time"
)

type Num struct {
	i int
	j int
}

func NewNum() *Num {
	return &Num{}
}

func getfun(n *Num) func() {
	return func() {
		time.Sleep(1 * time.Second)
		n.i, n.j = n.j, n.i
		fmt.Println(n)
	}
}

func main() {
	pool := grpool.New(100)

	for i := 1; i < 101; i++ {
		n := NewNum()
		n.i = i
		n.j = i + 1
		f := getfun(n)
		pool.Add(f)
	}
	fmt.Println("worker:", pool.Size())
	fmt.Println("  jobs:", pool.Jobs())
	gtime.SetInterval(time.Second, func() bool {
		fmt.Println("worker:", pool.Size())
		fmt.Println("  jobs:", pool.Jobs())
		fmt.Println()
		return true
	})

	select {}
}
