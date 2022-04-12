package main

import (
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	req := requests.Requests()
	resp, _ := req.Get(
		"http://localhost:8888",
		false,
		requests.Auth{"zhangdapeng520", "password...."},
	)
	println(resp.Text())
}
