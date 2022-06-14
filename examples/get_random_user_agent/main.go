package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		Debug: true,
	})

	i := 0
	for i < 10000 {
		fmt.Println(r.GetRandomUserAgent())
		i++
	}
}
