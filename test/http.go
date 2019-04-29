package main

import (
	"fmt"
	"godial/server"
	//"io/ioutil"
	"net"
	//"net/http"
	//"time"
)

func main() {
	/*
		c := &http.Client{
			Timeout: 2 * time.Second,
		}
		t1 := time.Now().UnixNano()
		//resp, err := c.Get("https://114.80.174.47")
		resp, err := c.Get("http://192.168.38.2")
		//resp, err := c.Get("http://54.95.242.111:443")
		//resp, err := c.Get("http://192.168.6.190:13636")
		if err != nil {
			// handle error
			fmt.Println("get error timeout")
			return
		}
		t2 := time.Now().UnixNano()

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			fmt.Println("read error timeout")
			return
		}

		fmt.Println(string(body))
		fmt.Println("ContentLength=", resp.ContentLength, len(string(body)))
		fmt.Println("delay=", (t2-t1)/1000)

		t3 := (t2 - t1) / 1000
		size := len(string(body))
		s := size * 1000 * 1000 / int(t3) / 1024
		fmt.Println("size=", size, "t=", t3, "speed=", s)

		//s1 := (float32(size) / 1024) / (float32(t3) / 1000 / 1000)
		//s1:=(size*1000*1000*1.0/total_time)/(1024*1024)
		s1 := (float32(len(string(body))) / 1024) / (float32(t3) / 1000 / 1000)
		fmt.Println(s1)
		fmt.Println("status_code=", resp.StatusCode)
		fmt.Println("url=", resp.Request.URL)
		fmt.Println("url=", resp.Request.URL.Scheme, resp.Request.URL.Host)
	*/
	fmt.Println(server.HttpDial(net.ParseIP("192.168.6.190")))
	fmt.Println(server.HttpDial(net.ParseIP("192.168.38.2")))
}
