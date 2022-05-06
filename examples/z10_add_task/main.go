package main

import (
	"fmt"
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
	var data zdpgo_requests.JsonData = map[string]interface{}{
		"taskName":              "测试任务1",
		"basStrategyTemplateId": 2,
		"equipmentIdList":       []int{27},
		"timeMode":              1,
		"xmazeNodeId":           "rOiRLKpTkmjjMddz",
		"targetType":            0,
		"taskPattern":           1,
		"simulateTargetIds":     []int{6},
		"responseTimeout":       12000,
		"agentHelperIdList":     []int{7}}

	// 发送请求
	resp, err := r.PostIgnoreParseError(url, data)
	if err != nil {
		fmt.Println("错误2", err)
	}
	fmt.Println(resp.Text())
}
