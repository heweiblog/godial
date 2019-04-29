package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

/*
#include <sys/resource.h>
int set_file_limit()
{
		struct rlimit tmp = {262143,262144};
		int rtn = setrlimit(RLIMIT_NOFILE,&tmp);
		if(rtn != 0)
		{
				return -1;
		}
		return 0;
}
*/
import "C"

func test(i int) {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	t1 := time.Now()
	fmt.Println(t1)
	//request := "http://" + "192.168.13.30" + ":80"
	request := "http://" + "192.168.13." + strconv.Itoa(i) + ":80"
	resp, err := c.Get(request)
	if err != nil {
		fmt.Println(err)
		t2 := time.Now()
		fmt.Println(t2.Sub(t1), "timeout")
	} else {
		t2 := time.Now()
		fmt.Println(t2)
		fmt.Println("192.168.13."+strconv.Itoa(i), int32((t2.Sub(t1).Nanoseconds())/1000), true)
		defer resp.Body.Close()
	}

}

func main() {
	//C.set_file_limit()
	for j := 0; j < 5; j++ {
		for i := 3; i < 100; i++ {
			go test(i)
		}
	}
	select {}
}
