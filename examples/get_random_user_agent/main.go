package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New(zdpgo_log.NewWithDebug(true, "log.log"))

	i := 0
	for i < 10000 {
		fmt.Println(r.GetRandomUserAgent())
		i++
	}
}
