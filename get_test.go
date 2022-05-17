package zdpgo_requests

import (
	"fmt"
	"testing"
)

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
