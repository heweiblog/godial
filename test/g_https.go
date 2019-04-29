package main

import (
	"fmt"
	"gitee.com/johng/gf/g/net/ghttp"
	"time"
)

func main() {
	t_start := time.Now()
	c := ghttp.NewClient()
	r, e := c.Get("http://192.168.13.94:80")
	if e != nil {
		fmt.Println("get error\n")
		return
	}
	s := r.ReadAll()
	fmt.Println(string(s))
	r.Close()
	t_end := time.Now()

	dur := t_end.Nanosecond() - t_start.Nanosecond()
	fmt.Println(t_end, t_start)
	fmt.Println(t_end.Nanosecond(), t_start.Nanosecond())
	fmt.Println(t_end.Sub(t_start).Nanoseconds() / 1000)
	fmt.Println("time_sub:", dur)
}
