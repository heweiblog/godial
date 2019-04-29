package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	/*
		        client与server进行通信时 client也要对server返回数字证书进行校验
				        因为server自签证书是无效的 为了client与server正常通信
						        通过设置客户端跳过证书校验
								        TLSClientConfig:{&tls.Config{InsecureSkipVerify: true}
										        true:跳过证书校验
	*/
	//tr.Dial("tcp", "114.80.184.124:443")

	client := &http.Client{
		Timeout:   2 * time.Second,
		Transport: tr}
	//resp, err := client.Get("https://115.239.210.27")
	//resp, err := client.Get("https://14.215.178.60")
	//resp, err := client.Get("https://114.80.184.124")
	//resp, err := client.Get("https://www.taobao.com/")
	//resp, err := client.Get("https://54.95.242.111:443/")
	//resp, err := client.Get("http://54.95.242.111:80/")
	resp, err := client.Get("https://14.18.240.22")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(resp)
}
