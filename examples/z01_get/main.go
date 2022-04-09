package main

import "github.com/zhangdapeng520/zdpgo_requests/libs/requests"

func main() {
	resp, err := requests.Get("http://www.zhanluejia.net.cn")
	if err != nil {
		return
	}
	println(resp.Text())
}
