package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/libs/requests"
)

func demo1() {
	r := zdpgo_requests.New()
	data := map[string]string{
		"name": "request123",
	}
	resp, _ := r.Post("http://localhost:8888", data)
	println(resp.Text())
}

func demo2() {
	r := zdpgo_requests.New()

	// 会被当做表单数据传递
	var data requests.Datas = map[string]string{
		"name": "requests_post_test",
	}
	resp, _ := r.Post("http://localhost:8888", data)
	println(resp.Text())
}

func demo3String() {
	r := zdpgo_requests.New()

	// 会被当做表单数据传递
	var data = "abc123张大鹏"
	resp, _ := r.Post("http://localhost:8888", data)
	println(resp.Text())
}

func main() {
	//demo1()
	//demo2()
	demo3String()
}
