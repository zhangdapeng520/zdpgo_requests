package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	// 直接设置请求头
	req := zdpgo_requests.New()
	req.Request.Header.Set("accept-encoding", "gzip, deflate, br")
	resp, _ := req.Get("http://localhost:8888", false, requests.Header{"Referer": "http://127.0.0.1:9999"})
	println(resp.Text())

	// 将请求头作为参数传递
	h := requests.Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}
	h2 := requests.Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"User-Agent":      "zdpgo_requests",
	}
	resp, _ = req.Get("http://localhost:8888", h, h2)
	println(resp.Text())
}
