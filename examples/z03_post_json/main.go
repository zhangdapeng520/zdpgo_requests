package main

import "github.com/zhangdapeng520/zdpgo_requests/libs/requests"

func main() {
	jsonStr := "{\"name\":\"requests_post_test\"}"
	resp, _ := requests.PostJson("https://www.httpbin.org/post", jsonStr)
	println(resp.Text())
}
