package zdpgo_requests

import (
	"fmt"
	"testing"
)

func getRequests() *Requests {
	return NewWithConfig(Config{
		Debug: true,
	})
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
	var jsonStr = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ = r.Post(url, true, jsonStr, headers)
	println(resp.Text())

	// 权限校验
	resp, _ = r.Get(
		url,
		BaseAuth{"zhangdapeng520", "password...."},
	)
	println(resp.Text())
}

// 测试设置请求头
func TestRequests_header(t *testing.T) {
	// 直接设置请求头
	req := getRequests()
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	resp, _ := req.Get("http://localhost:8888", false, Header{"Referer": "http://127.0.0.1:9999"})
	println(resp.Text())

	// 将请求头作为参数传递
	h := Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}
	h2 := Header{
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
	p := Param{
		"name": "file",
		"id":   "12345",
	}
	resp, _ := req.Get("http://localhost:8888", p)
	fmt.Println(resp.Text())
}

func TestRequests_redirect(t *testing.T) {
	req := getRequests()
	url := "https://www.baidu.com/link?url=IIZcBDQ9FSkK8wRluFkNAxjf4a7VDwHH0kFqGazjEAFGRDdnxe0HqQRdSocksxbbrpMjo7PTBeGjgnmf0aYOqN7ld6dXDBVO_jMYS16Yuy7CI5M_TMysMLpmFhF4CEjGjXOEYvjL_r9Hgz2-4jwsoa"
	resp, _ := req.Get(url)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.IsRedirect)
	fmt.Println(resp.RedirectUrl)
	fmt.Println(resp.StartTime)
	fmt.Println(resp.EndTime)
	fmt.Println(resp.SpendTime)
	fmt.Println(resp.SpendTimeSeconds)
}

// 测试代理的使用
func TestRequests_proxy(t *testing.T) {
	req := getRequests()

	// 设置代理
	err := req.SetProxy("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	// 发送请求
	// 设置了代理以后，请求被重定向了代理的URL
	resp, _ := req.Get("http://127.0.0.1:3333/ping")
	fmt.Println("响应：", resp.Text())
}

// 测试删除文件夹
func TestRequests_DeleteDir(t *testing.T) {
	r := getRequests()

	fmt.Println(r.Config.FsTmpDir)
	fmt.Println(r.Exists(r.Config.FsTmpDir))
	r.DeleteDir(r.Config.FsTmpDir)
	fmt.Println(r.Exists(r.Config.FsTmpDir))
}
