package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_requests.New()
	host := "10.1.4.2:8080"
	baseUrl := "http://" + host + "/pte/api-v1/bas/task/add"
	query := ""
	url := baseUrl + query

	// 数据
	Json := zdpgo_json.New()
	var data zdpgo_requests.JsonData
	Json.Load("examples/poc/data.json", &data)

	// 发送请求
	resp, err := r.PostIgnoreParseError(url, data)
	if err != nil {
		fmt.Println("错误2", err)
	}
	fmt.Println(resp.Text())
}
