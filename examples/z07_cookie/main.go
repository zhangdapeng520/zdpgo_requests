package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests/libs/requests"
)

func main() {
	req := requests.Requests()
	resp, _ := req.Get("https://www.httpbin.org")
	coo := resp.Cookies()
	// coo is [] *http.Cookies
	println("********cookies*******")
	for _, c := range coo {
		fmt.Println(c.Name, c.Value)
	}
}
