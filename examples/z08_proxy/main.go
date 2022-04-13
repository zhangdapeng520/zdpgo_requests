package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	req := requests.Requests()

	// 设置代理
	req.Proxy("http://localhost:8888")

	// 发送请求
	// 设置了代理以后，请求被重定向了代理的URL
	resp, _ := req.Get("http://localhost:9999", false)

	fmt.Println("响应：", resp.Text())
}
