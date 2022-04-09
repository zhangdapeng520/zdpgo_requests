package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests/libs/requests"
)

func main() {
	req := requests.Requests()
	p := requests.Params{
		"title": "The blog",
		"name":  "file",
		"id":    "12345",
	}
	resp, _ := req.Get("http://www.cpython.org", p)
	fmt.Println(resp.Text())
}
