package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_requests.New()
	url := "http://localhost:8888"
	resp, err := r.Get(url, true)
	if err != nil {
		return
	}
	println(resp.Text())

	// 发送POST请求
	data := map[string]string{
		"name": "request123",
	}
	resp, _ = r.Post(url, data)
	println(resp.Text())

	// 发送json数据
	var jsonStr requests.Datas = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers requests.Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ = requests.Post(url, true, jsonStr, headers)
	println(resp.Text())

	// 权限校验
	resp, _ = r.Get(
		url,
		requests.Auth{"zhangdapeng520", "password...."},
	)
	println(resp.Text())
}
