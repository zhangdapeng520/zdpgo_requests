package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	req := zdpgo_New()
	resp, _ := req.Get(
		"http://localhost:8080/admin/secrets",
		zdpgo_Auth{"zhangdapeng", "zhangdapeng"},
	)
	println(resp.Text())
}
