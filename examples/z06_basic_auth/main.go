package main

import (
	"github.com/zhangdapeng520/zdpgo_requests/libs/requests"
)

func main() {
	req := requests.Requests()
	resp, _ := req.Get("https://api.github.com/user", requests.Auth{"asmcos", "password...."})
	println(resp.Text())
}
