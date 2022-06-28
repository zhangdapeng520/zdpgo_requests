package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
	"reflect"
)

func main() {
	requests := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{})
	target := "http://127.0.0.1:3333/aes"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"

	// 同时有10w个并发
	for i := 0; i < 100000; i++ {
		response, err := requests.PostAes(target, jsonStr)
		if err != nil {
			panic(err)
		}
		fmt.Println(response.Text, reflect.TypeOf(response.Text))
	}
}
