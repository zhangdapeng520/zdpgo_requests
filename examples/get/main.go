package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New(zdpgo_log.NewWithDebug(true, "log.log"))

	response, err := r.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Text)
}
