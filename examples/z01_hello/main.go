package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_requests.New()
	url := "http://localhost:8888/payload/xxx"
	resp, err := r.Get(url, true)
	if err != nil {
		return
	}
	println(resp.StatusCode, resp.Text())

	// 发送POST请求
	data := map[string]string{
		"name": "request123",
	}
	resp, _ = r.Post(url, data)
	println(resp.StatusCode, resp.Text())

	// 发送json数据
	var jsonStr zdpgo_requests.Datas = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers zdpgo_requests.Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ = zdpgo_requests.Post(url, true, jsonStr, headers)
	println(resp.StatusCode, resp.Text())

	// 权限校验
	resp, _ = r.Get(
		url,
		zdpgo_requests.Auth{"zhangdapeng520", "password...."},
	)
	println(resp.StatusCode, resp.Text())
}
