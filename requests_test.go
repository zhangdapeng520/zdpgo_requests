package zdpgo_requests

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
	"testing"
)

func getRequests() *Requests {
	return New()
}

// 测试基本使用
func TestRequests_basic(t *testing.T) {
	// 发送GET请求
	r := getRequests()
	url := "http://localhost:8888"
	resp, err := r.Get(url, true)
	if err != nil {
		return
	}
	println(resp.Text())

	// 发送POST请求
	data := map[string]string{
		"name": "request123",
	}
	resp, _ = r.Post(url, data)
	println(resp.Text())

	// 发送json数据
	var jsonStr requests.Datas = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers requests.Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ = requests.Post(url, true, jsonStr, headers)
	println(resp.Text())

	// 权限校验
	resp, _ = r.Get(
		url,
		requests.Auth{"zhangdapeng520", "password...."},
	)
	println(resp.Text())
}

// 测试设置请求头
func TestRequests_header(t *testing.T) {
	// 直接设置请求头
	req := getRequests()
	req.Request.Header.Set("accept-encoding", "gzip, deflate, br")
	resp, _ := req.Get("http://localhost:8888", false, requests.Header{"Referer": "http://127.0.0.1:9999"})
	println(resp.Text())

	// 将请求头作为参数传递
	h := requests.Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}
	h2 := requests.Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"User-Agent":      "zdpgo_requests",
	}
	resp, _ = req.Get("http://localhost:8888", h, h2)
	println(resp.Text())
}

// 测试设置查询参数
func TestRequests_params(t *testing.T) {
	req := getRequests()
	p := requests.Params{
		"name": "file",
		"id":   "12345",
	}
	resp, _ := req.Get("http://localhost:8888", false, p)
	fmt.Println(resp.Text())
}
