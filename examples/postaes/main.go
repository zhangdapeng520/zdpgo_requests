package main

import (
	"fmt"
	"reflect"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	requests := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		Debug: true,
	})
	target := "http://127.0.0.1:3333/aes"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"
	response, err := requests.PostAes(target, jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text, reflect.TypeOf(response.Text))
}