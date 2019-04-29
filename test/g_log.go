package main

import (
	//"gitee.com/johng/gf/g"
	//"gitee.com/johng/gf/g/os/gfile"
	"gitee.com/johng/gf/g/container/gtype"
	"gitee.com/johng/gf/g/os/glog"
	"time"
)

var Num *gtype.Int

func test1() {
	for i := 0; i < 5; i++ {
		Num.Add(1)
		glog.Println("111", Num.Val())
	}
}
func test2() {
	for i := 0; i < 5; i++ {
		Num.Add(1)
		glog.Println("222", Num.Val())
	}
}
func test3() {
	for i := 0; i < 5; i++ {
		Num.Add(1)
		glog.Println("333", Num.Val())
	}
}

// 设置日志等级
func main() {
	path := "./glog"
	glog.SetPath(path)
	glog.SetStdPrint(false)
	// 使用默认文件名称格式
	glog.Println("标准文件名称格式，使用当前时间时期")
	glog.Println("标准文件名称格式，使用当前时间时期")
	Num = gtype.NewInt()
	go test1()
	go test2()
	go test3()
	time.Sleep(time.Second)
	// 通过SetFile设置文件名称格式
	/*
		glog.SetFile("stdout.log")
		glog.Println("设置日志输出文件名称格式为同一个文件")
		// 链式操作设置文件名称格式
		glog.File("stderr.log").Println("支持链式操作")
		glog.File("error-{Ymd}.log").Println("文件名称支持带gtime日期格式")
		glog.File("access-{Ymd}.log").Println("文件名称支持带gtime日期格式")

		list, err := gfile.ScanDir(path, "*")
		g.Dump(err)
		g.Dump(list)
	*/
}
