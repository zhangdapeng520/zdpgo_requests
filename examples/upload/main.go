package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New()

	uploadUrl := "http://localhost:3333/upload"

	// 上传文件
	response, err := r.Upload(uploadUrl, "file", "README.md")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Text)

	// 上传字节数组
	response, err = r.UploadByBytes(uploadUrl, "file", "hello.txt", []byte("你好，世界~"))
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Text)
}
