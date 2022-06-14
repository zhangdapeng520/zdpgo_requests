package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		Debug: true,
	})

	response, err := r.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Text)
}
