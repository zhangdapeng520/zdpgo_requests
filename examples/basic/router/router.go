package router

import (
	"embed"
	"github.com/zhangdapeng520/zdpgo_requests"
)

/*
@Time : 2022/5/20 9:50
@Author : 张大鹏
@File : router.go
@Software: Goland2021.3.1
@Description: 测试路由相关的用法
*/

func Any(r *zdpgo_requests.Requests, url string) {
	resp, err := r.Any("get", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("post", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("put", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("delete", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("patch", url, true)
	println(resp.Text(), err)
}

func Special(r *zdpgo_requests.Requests, url string) {
	resp, err := r.Get(url)
	println(resp.Text(), err)
	resp, err = r.Post(url)
	println(resp.Text(), err)
	resp, err = r.Put(url)
	println(resp.Text(), err)
	resp, err = r.Delete(url)
	println(resp.Text(), err)
	resp, err = r.Patch(url)
	println(resp.Text(), err)
}

func Proxy(r *zdpgo_requests.Requests, url, proxyUrl string) {
	r.SetProxy(proxyUrl)
	resp, err := r.Get(url)
	println(resp.Text(), err)
	resp, err = r.Post(url)
	println(resp.Text(), err)
	resp, err = r.Put(url)
	println(resp.Text(), err)
	resp, err = r.Delete(url)
	println(resp.Text(), err)
	resp, err = r.Patch(url)
	println(resp.Text(), err)
	r.RemoveProxy()
}
func Params(r *zdpgo_requests.Requests, url string) {
	param := zdpgo_requests.Param{
		"a": "11",
		"b": "22.222",
		"c": "abc",
	}
	resp, err := r.Get(url, param)
	println(resp.Text(), err)
	resp, err = r.Post(url, param)
	println(resp.Text(), err)
	resp, err = r.Put(url, param)
	println(resp.Text(), err)
	resp, err = r.Delete(url, param)
	println(resp.Text(), err)
	resp, err = r.Patch(url, param)
	println(resp.Text(), err)
}

// UploadFS 上传fs内嵌系统文件
func UploadFS(r *zdpgo_requests.Requests, url string, fsObj embed.FS) {
	r.IsFs = true
	r.Fs = fsObj
	fileMap := zdpgo_requests.Files{
		"file": "test/test.txt",
	}
	resp, err := r.Post(url, fileMap)
	println(resp.Text(), err)
}

func Timeout(r *zdpgo_requests.Requests, url string) {
	resp, err := r.Get(url)
	println(resp.Text(), err)
}

func Auth(r *zdpgo_requests.Requests, url string, username, password string) {
	r.SetBasicAuth(username, password)
	resp, err := r.Get(url)
	println(resp.Text(), err)
}
