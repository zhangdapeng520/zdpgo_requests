package zdpgo_requests

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func (r *Requests) GetHttpRequest(request Request) *http.Request {
	// 请求方法
	if request.Method == "" {
		request.Method = "GET"
	}

	// 请求头
	header := make(http.Header)
	if request.Header != nil {
		for k, v := range request.Header {
			header.Set(k, v)
		}
	}
	if header.Get("User-Agent") == "" {
		header.Set("User-Agent", r.Config.UserAgent)
	}

	// 构造请求对象
	req := &http.Request{
		Method:     request.Method,
		Header:     header,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	req.Close = true // 解决EOF的bug

	// 请求地址
	if request.Url != "" {
		// 查询参数
		urlPath := request.Url
		if request.Query != nil {
			var params []string
			for k, v := range request.Query {
				params = append(params, fmt.Sprintf("%s=%s", k, v))
			}
			queryStr := strings.Join(params, "&")
			if strings.Contains(urlPath, "?") && strings.Contains(urlPath, "=") {
				urlPath += "&" + queryStr
			} else if strings.Contains(urlPath, "?") {
				urlPath += queryStr
			} else {
				urlPath += "?" + queryStr
			}
		}

		// 请求地址
		urlPared, err := url.Parse(urlPath)
		if err != nil {
			r.Log.Error("解析URL失败", "err", err, "url", urlPath)
			return req
		}
		req.URL = urlPared
	}

	// 请求体
	if request.Body != nil {
		req.ContentLength = int64(request.Body.Len())
		buf := request.Body.Bytes()
		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(buf)
			return io.NopCloser(reader), nil
		}
		readCloser := io.NopCloser(request.Body)
		req.Body = readCloser
	}

	// 设置基础权限
	if request.BasicAuth.Username != "" && request.BasicAuth.Password != "" {
		req.SetBasicAuth(request.BasicAuth.Username, request.BasicAuth.Password)
	}

	// 返回
	return req
}

// GetHttpClient 获取HTTP请求的客户端
func (r *Requests) GetHttpClient() *http.Client {
	// 获取端口
	port := r.GetHttpPort()
	r.ClientPort = port

	// 绑定本地端口
	netAddr := &net.TCPAddr{Port: port}
	dialer := &net.Dialer{LocalAddr: netAddr}
	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          100, // 连接池大小
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: !r.Config.IsCheckHttps},
	}

	// 设置代理
	if r.Config.ProxyUrl != "" {
		uri, err := url.Parse(r.Config.ProxyUrl) // 解析代理地址
		if err != nil {
			r.Log.Error("解析代理地址失败", "error", err, "proxyUrl", r.Config.ProxyUrl)
		}
		tr.Proxy = http.ProxyURL(uri) // 设置代理
	}

	// 创建客户端
	httpClient := &http.Client{
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
	return httpClient
}

// GetHttpPort 获取系统中可用的端口号
func (r *Requests) GetHttpPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		r.Log.Error("解析TCP地址失败", "error", err)
		return 0
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		r.Log.Error("创建tcp监听失败", "error", err)
		return 0
	}
	defer l.Close()

	// 获取端口号
	p := l.Addr().(*net.TCPAddr).Port
	return p
}
