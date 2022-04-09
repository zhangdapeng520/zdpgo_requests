package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New()
	data := map[string]string{
		"name": "requests_post_test",
	}
	resp, _ := r.Post("https://www.httpbin.org/post", data)
	println(resp.Text())
}
