package main

import (
	"gitee.com/johng/gf/g/os/glog"
	"github.com/sevlyar/go-daemon"
	"godial/client"
	"godial/server"
	"log"
)

func log_init() {
	path := "/var/log/godial"
	glog.SetPath(path)
	glog.SetStdPrint(false)
}

func main() {

	cntxt := &daemon.Context{
		PidFileName: "godial_pid",
		PidFilePerm: 0644,
		LogFileName: "godial_daemon.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[godial]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	log.Print("- - - - - - - - - - - - - - -")
	log.Print("godial daemon started!!!")

	log_init()

	go client.Register()
	log.Print("Dial mouble register to ms")
	go server.Monitor()
	log.Print("Dial monnitor mouble start")
	log.Print("Dial thrift server start")
	server.Dial_server_start()
}
