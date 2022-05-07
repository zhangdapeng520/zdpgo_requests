package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	req := zdpgo_requests.New()
	p := zdpgo_requests.Params{
		"name": "file",
		"id":   "12345",
	}
	resp, _ := req.Get("http://localhost:8888", false, p)
	fmt.Println(resp.Text())
}
