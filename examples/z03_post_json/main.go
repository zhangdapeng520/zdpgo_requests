package main

import (
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

// 发送JSON字符串
func demo1() {
	jsonStr := "{\"name\":\"requests_post_test\"}"
	resp, _ := requests.PostJson("http://localhost:8888", jsonStr)
	println(resp.Text())
}

// 发送JSON字典
func demo2JsonMap() {
	var jsonStr requests.Datas = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers requests.Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ := requests.Post("http://localhost:8888", true, jsonStr, headers)
	println(resp.Text())
}

func main() {
	demo1()
	demo2JsonMap()
}
