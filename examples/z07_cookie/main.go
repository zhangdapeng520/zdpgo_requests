package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	req := zdpgo_requests.New()
	resp, _ := req.Get("http://localhost:8888", false)
	coo := resp.Cookies()
	println("********cookies*******")
	for _, c := range coo {
		fmt.Println(c.Name, c.Value)
	}
}
