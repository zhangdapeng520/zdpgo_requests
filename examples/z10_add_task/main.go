package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_requests.New()
	host := "10.1.3.12:8888"
	host = "10.1.4.2:8080"
	baseUrl := "http://" + host + "/pte/api-v1/open-api/bas/task/add"
	query := ""
	url := baseUrl + query

	// 数据
	var data requests.JsonData = map[string]interface{}{
		"taskName":              "测试1",
		"basStrategyTemplateId": 2,
		"equipmentId":           22,
		"timeMode":              1,
		"attackSevNodeId":       6,
		"verifyRulesWafType":    "sev_assistant",
	}

	// 发送请求
	resp, err := r.PostIgnoreParseError(url, data)
	if err != nil {
		fmt.Println("错误2", err)
	} else {
		println(resp.Text())
		println("请求详情：\n", resp.RawReqDetail)
		println("响应详情：\n", resp.RawRespDetail)
	}
}
