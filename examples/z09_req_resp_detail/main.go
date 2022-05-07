package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_requests.New()
	baseUrl := "http://10.1.3.12:8888/"
	query := "?a=<script>alert(\"XSS\");</script>&b=UNION SELECT ALL FROM information_schema AND ' or SLEEP(5) or '&c=../../../../etc/passwd"
	url := baseUrl + query

	var h1 zdpgo_requests.Header = zdpgo_requests.Header{"a": "111", "b": "222"}
	resp, err := r.GetIgnoreParseError(url, h1)
	if err != nil {
		fmt.Println("错误2", err)
	} else {
		println(resp.Text())
		println("请求详情：\n", resp.RawReqDetail)
		println("响应详情：\n", resp.RawRespDetail)
	}

	var h2 zdpgo_requests.Header = zdpgo_requests.Header{"c": "333", "d": "444"}
	resp1, err := r.GetIgnoreParseError(url, h2)
	if err != nil {
		fmt.Println("错误3", err)
	} else {
		println(resp1.Text())
		println("请求详情：\n", resp.RawReqDetail)
		println("响应详情：\n", resp.RawRespDetail)
	}
}
