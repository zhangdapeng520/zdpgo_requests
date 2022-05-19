package zdpgo_requests

import (
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
}
