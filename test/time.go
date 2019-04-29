package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
	time.Sleep(time.Second)
	fmt.Println("time.Second = ", time.Second)
	fmt.Println(time.Unix(1, 0))
	fmt.Println("time:", time.Now().Format("2006-01-02 15:04:05")) //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
}
