package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	requests := zdpgo_requests.New(zdpgo_log.NewWithDebug(true, "log.log"))
	target := "http://127.0.0.1:3333/aes"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"

	// 发送1w条请求
	i := 0
	for i < 100000 {
		response, err := requests.PostAes(target, jsonStr)
		if err != nil {
			panic(err)
		}
		fmt.Println(response.ToJsonStr())
		i++
	}
}
