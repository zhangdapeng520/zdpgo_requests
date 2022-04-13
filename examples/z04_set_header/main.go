package main

import (
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func demo1() {
	req := requests.Requests()

	resp, err := req.Get("http://localhost:8888", false, requests.Header{"Referer": "http://www.jeapedu.com"})
	if err == nil {
		println(resp.Text())
	}
}

func demo2() {
	req := requests.Requests()
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	resp, _ := req.Get("http://localhost:8888", false, requests.Header{"Referer": "http://www.jeapedu.com"})
	println(resp.Text())
}

func demo3() {
	req := requests.Requests()
	h := requests.Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}

	h2 := requests.Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"User-Agent":      "zdpgo_requests",
	}
	resp, _ := req.Get("http://localhost:8888", false, h, h2)
	println(resp.Text())
}

func main() {
	demo1()
	demo2()
	demo3()
}