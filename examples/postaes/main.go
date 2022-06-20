package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_requests"
	"reflect"
)

func main() {
	requests := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{}, zdpgo_log.Tmp)
	target := "http://127.0.0.1:3333/aes"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"
	response, err := requests.PostAes(target, jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text, reflect.TypeOf(response.Text))
}
