package zdpgo_requests

import (
	"fmt"
	"testing"
)

// 测试代理的使用
func TestRequests_proxy(t *testing.T) {
	req := getRequests()

	// 设置代理
	err := req.SetProxy("http://localhost:8888")
	if err != nil {
		panic(err)
	}

	// 发送请求
	// 设置了代理以后，请求被重定向了代理的URL
	resp, _ := req.Get("http://localhost:9999", false)
	fmt.Println("响应：", resp.Text())
}
