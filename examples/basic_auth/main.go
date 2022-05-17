package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	req := zdpgo_requests.New()
	resp, _ := req.Post(
		"http://localhost:3333/admin",
		zdpgo_requests.BaseAuth{Username: "zhangdapeng", Password: "zhangdapeng"},
	)
	println(resp.Text())
}
