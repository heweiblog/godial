package main

import (
	"fmt"
	"gitee.com/johng/gf/g/container/gtype"
)

func main() {
	// 创建一个Int型的并发安全基本类型对象
	i := gtype.NewInt()

	// 设置值
	fmt.Println(i.Set(10))

	// 获取值
	fmt.Println(i.Val())

	// 数值1，并返回修改之后的数值
	fmt.Println(i.Add(1))
	fmt.Println(i.Add(1))
	fmt.Println(i.Add(1))
	fmt.Println(i.Add(1))
}
