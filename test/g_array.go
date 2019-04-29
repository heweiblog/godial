package main

import (
	"fmt"
	"gitee.com/johng/gf/g/container/garray"
)

func main() {
	arr := garray.NewArray(10, 20, true)

	fmt.Println(arr.Len())

	arr.Append(100)
	arr.Append(110)
	arr.Append("hww")

	fmt.Println(arr.Len())

}
