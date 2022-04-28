package main

import (
	"embed"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

//go:embed test/*
var fsObj embed.FS

func main() {
	r := zdpgo_New()

	targetUrl := "http://localhost:8888/upload"
	filename := "test/main.go"

	respBytes, err := r.UploadFsToBytes(targetUrl, fsObj, "file", filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(respBytes))

	respString, err := r.UploadFsToString(targetUrl, fsObj, "file", filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(respString)
}
