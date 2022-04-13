package zdpgo_requests

import "testing"

// 测试各种HTTP请求方法
func TestRequests_method(t *testing.T) {
	r := getRequests()
	targetUrl := "http://localhost:8888"

	// 发送GET请求
	resp, _ := r.Get(targetUrl)
	println(resp.Text())

	// 发送POST请求
	resp, _ = r.Post(targetUrl)
	println(resp.Text())

	// 发送PUT请求
	resp, _ = r.Put(targetUrl)
	println(resp.Text())

	// 发送DELETE请求
	resp, _ = r.Delete(targetUrl)
	println(resp.Text())

	// 发送PATCH请求
	resp, _ = r.Patch(targetUrl)
	println(resp.Text())
}
