package zdpgo_requests

import (
	"embed"
	"fmt"
	"testing"
)

/*
@Time : 2022/5/19 14:55
@Author : 张大鹏
@File : method_test.go
@Software: Goland2021.3.1
@Description: method相关测试
*/

//go:embed tmp
var fsObj embed.FS

// 测试各种HTTP请求方法
func TestRequests_Get(t *testing.T) {
	r := getRequests()
	targetUrl := "http://localhost:3333/ping?a=111&b=222#abc"

	resp, _ := r.Get(targetUrl)
	println(resp.Text())

	resp, _ = r.Get("https://bas.x-check.cn//upload/RjVzTFNBltCRfYbs")
	println(resp.Text())

	resp, _ = r.Post(targetUrl)
	println(resp.Text())

	resp, _ = r.Put(targetUrl)
	println(resp.Text())

	resp, _ = r.Delete(targetUrl)
	println(resp.Text())

	resp, _ = r.Patch(targetUrl)
	println(resp.Text())

	// 权限
	resp, _ = r.Get("http://localhost:3333/admin", BaseAuth{"zhangdapeng", "zhangdapeng"})
	println(resp.Text())
}

// 测试请求https
func TestRequests_GetHttps(t *testing.T) {
	r := getRequests()

	resp, err := r.Get("https://bas.x-check.cn//upload/RjVzTFNBltCRfYbs")
	if err != nil {
		fmt.Println("请求数据失败", err)
	}
	println(resp.Text())
}

// 测试任意方法
func TestRequests_Any(t *testing.T) {
	r := getRequests()
	var (
		resp *Response
		err  error
		url  = "http://localhost:8889/payload/"
	)

	// any方法
	resp, err = r.Any("get", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("post", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("put", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("delete", url, true)
	println(resp.Text(), err)
	resp, err = r.Any("patch", url, true)
	println(resp.Text(), err)

	// 特定方法
	resp, err = r.Get(url)
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

func TestRequests_Upload1(t *testing.T) {
	r := getRequests()
	url := "http://localhost:8889/upload"
	// 上传文件
	r.Files = append(r.Files, map[string]string{
		"file": "tmp/img1.jpg",
	})
	fmt.Println("================")
	resp, err := r.Post(url)
	println(resp.Text(), err)
}

// 测试上传嵌入文件系统
func TestRequests_UploadFS(t *testing.T) {
	r := getRequests()

	// 设置代理
	r.SetProxy("http://localhost:8080")

	url := "http://localhost:8889/upload"

	// 上传文件
	r.Files = append(r.Files, map[string]string{
		"file": "tmp/img1.jpg",
	})

	// 指定使用嵌入文件系统
	r.IsFs = true
	r.Fs = fsObj

	// 执行上传
	resp, err := r.Post(url)
	println(resp.Text(), err)
}
