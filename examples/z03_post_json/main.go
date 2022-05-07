package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New()

	// 发送JSON字符串
	var jsonStr zdpgo_requests.JsonString = "{\"name\":\"requests_post_test\"}"
	resp, _ := r.Post("http://localhost:8888", jsonStr)
	println(resp.Text())

	// 发送map
	var data zdpgo_requests.JsonData = make(map[string]interface{})
	data["name"] = "root"
	data["password"] = "root"
	data["host"] = "localhost"
	data["port"] = 8888
	resp, _ = r.Post("http://localhost:8888", data)
	println(resp.Text())

	// PUT发送JSON数据
	resp, _ = r.Put("http://localhost:8888", data)
	println(resp.Text())

	// DELETE发送JSON数据
	resp, _ = r.Delete("http://localhost:8888", data)
	println(resp.Text())

	// PATCH发送JSON数据
	resp, _ = r.Patch("http://localhost:8888", data)
	println(resp.Text())

	// 发送纯文本数据（非json）
	resp, _ = r.Post("http://localhost:8888", "纯文本内容。。。")
	println(resp.Text())
}
