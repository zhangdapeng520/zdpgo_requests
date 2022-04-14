package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	req := zdpgo_requests.New()
	resp, _ := req.Get(
		"http://localhost:8080/admin/secrets",
		requests.Auth{"zhangdapeng", "zhangdapeng"},
	)
	println(resp.Text())
}
