package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_New()
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	err := r.Download(imgUrl, "test1.jpg")
	if err != nil {
		panic(err)
	}
}
