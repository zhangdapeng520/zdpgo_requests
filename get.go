package zdpgo_requests

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func (r *Requests) GetHttpRequest() (req *http.Request) {
	req = &http.Request{
		Method:     "GET",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	// 设置请求头
	req.Header.Set("Content-Type", r.Config.ContentType)
	req.Header.Set("User-Agent", r.Config.UserAgent)

	// 返回
	return
}

// GetHttpClient 获取HTTP请求的客户端
func (r *Requests) GetHttpClient() (httpClient *http.Client) {
	// 是否跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.CheckHttps},
	}

	// 创建客户端
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(r.Config.Timeout),
	}

	// 超时控制
	if r.Config.Timeout != 0 {
		httpClient.Timeout = time.Duration(r.Config.Timeout) * time.Second
	}

	// 自动生成cookie
	jar, err := cookiejar.New(nil)
	if err != nil {
		r.Log.Error("创建cookie失败", "error", err)
	}
	httpClient.Jar = jar

	// 返回
	return
}
