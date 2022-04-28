package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_New()
	imgUrl := "http://localhost:8888/upload"
	err := r.Upload(imgUrl, "test1.jpg")
	if err != nil {
		panic(err)
	}
}
