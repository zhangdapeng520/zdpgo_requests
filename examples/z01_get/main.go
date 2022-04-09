package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New()
	resp, err := r.Get("http://www.zhanluejia.net.cn")
	if err != nil {
		return
	}
	println(resp.Text())
}
