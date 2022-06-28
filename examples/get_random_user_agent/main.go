package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New()

	i := 0
	for i < 10000 {
		fmt.Println(r.GetRandomUserAgent())
		i++
	}
}
