package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	req := requests.Requests()
	p := requests.Params{
		"title": "The blog",
		"name":  "file",
		"id":    "12345",
	}
	resp, _ := req.Get("http://localhost:8888", false, p)
	fmt.Println(resp.Text())
}
