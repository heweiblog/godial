package server

import (
	"crypto/tls"
	"gitee.com/johng/gf/g/os/glog"
	"godial/gen-go/rpc/yamutech/com"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func HttpDial(a net.IP) (int32, bool) {
	c := &http.Client{
		Timeout: 2 * time.Second,
	}

	request := "http://" + a.String()
	t := time.Now()
	resp, err := c.Get(request)
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()

	return int32(time.Since(t).Nanoseconds() / 1000), true
}

func WebDial(dname string, a net.IP, ip_result *com.IpResult_) {

	c := &http.Client{
		Timeout: 2 * time.Second,
	}
	t_start := time.Now()
	request := "http://" + a.String() + ":80"
	resp, err := c.Get(request)
	if err != nil {
		return
	}
	t_end := time.Now()

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == 200 {
		ip_result.Delay = int32(t_end.Sub(t_start).Nanoseconds() / 1000)
		ip_result.Available = true

		vres := com.NewVideoResult_()
		vres.Speed = int32(len(string(body)) * 1000 * 1000 / int(ip_result.Delay) / 1024) //kb/s
		vres.Available = true
		vres.URL = dname

		ip_result.VideoResults = append(ip_result.VideoResults, vres)
	}
}

func HttpHttpsDial(dname string, a net.IP, ip_result *com.IpResult_) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   2 * time.Second,
		Transport: tr}
	request := "http://" + a.String()
	t_start := time.Now()
	resp, err := client.Get(request)
	if err != nil {
		glog.Println("error:", request, err)
		return
	}
	t_end := time.Now()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//glog.Println("HTTP/HTTPS:", request, len(string(body)))
	if resp.StatusCode == 200 {
		ip_result.Delay = int32(t_end.Sub(t_start).Nanoseconds() / 1000)
		ip_result.Available = true
		ip_result.Downloadspeed = int32(len(string(body)) * 1000 * 1000 / int(ip_result.Delay) / 1024) //kb/s

		vres := com.NewVideoResult_()
		vres.Speed = int32(len(string(body)) * 1000 * 1000 / int(ip_result.Delay) / 1024) //kb/s
		vres.Available = true
		vres.URL = dname
		ip_result.VideoResults = append(ip_result.VideoResults, vres)

		return
	}

	request = "https://" + a.String()
	t_start = time.Now()
	res, er := client.Get(request)
	if er != nil {
		return
	}
	t_end = time.Now()
	defer res.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if res.StatusCode == 200 {
		ip_result.Delay = int32(t_end.Sub(t_start).Nanoseconds() / 1000)
		ip_result.Available = true
		ip_result.Downloadspeed = int32(len(string(body)) * 1000 * 1000 / int(ip_result.Delay) / 1024) //kb/s

		vres := com.NewVideoResult_()
		vres.Speed = int32(len(string(body)) * 1000 * 1000 / int(ip_result.Delay) / 1024) //kb/s
		vres.Available = true
		vres.URL = dname
		ip_result.VideoResults = append(ip_result.VideoResults, vres)

		return
	}
}
