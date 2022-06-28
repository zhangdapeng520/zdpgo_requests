package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	requests := zdpgo_requests.New()
	target := "http://127.0.0.1:3333/aes"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"

	// 发送1w条请求
	i := 0
	for i < 100000 {
		response, err := requests.PostAes(target, jsonStr)
		if err != nil {
			requests.Log.Error("解密AES数据失败", "error", err)
		}
		fmt.Println(response.ToJsonStr())
		i++
	}
}
