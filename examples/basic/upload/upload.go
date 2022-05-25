package upload

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"io/ioutil"
)

/*
@Time : 2022/5/25 20:32
@Author : 张大鹏
@File : upload.go
@Software: Goland2021.3.1
@Description:
*/

func Upload(r *zdpgo_requests.Requests) {
	url := "http://localhost:8888/upload"
	r.Upload(url, "file", "test1.txt")
	r.Log.Debug("响应结果", "response", r.Response.Text)
}

func UploadByBytes(r *zdpgo_requests.Requests) {
	url := "http://localhost:8888/upload"
	fileContent, _ := ioutil.ReadFile("test1.txt")
	r.UploadByBytes(url, "file", "test1.txt", fileContent)
	r.Log.Debug("响应结果", "response", r.Response.Text)
}
