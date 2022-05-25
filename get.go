package zdpgo_requests

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
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
		urlPared, err := url.Parse(request.Url)
		if err != nil {
			r.Log.Error("解析URL失败", "err", err, "url", request.Url)
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

	// 返回
	return req
}

// GetHttpClient 获取HTTP请求的客户端
func (r *Requests) GetHttpClient() *http.Client {
	// 获取端口
	port := r.GetHttpPort()
	r.ClientPort = port
	r.Response = &Response{
		ClientPort: port,
	}

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
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: !r.Config.CheckHttps},
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

// GetParsedUrl 获取解析后的URL地址
func (r *Requests) GetParsedUrl(userURL string) string {
	// 解析URL
	parsedURL, err := url.Parse(userURL)
	if err != nil {
		r.Log.Error("解析URL地址失败", "error", err, "userURL", userURL)
		return userURL
	}

	// 解析Query查询参数
	if parsedURL != nil {
		parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
		if err != nil {
			r.Log.Error("解析query查询参数失败", "error", err, "query", parsedURL.RawQuery)
			// 无法正常解析query参数，尝试将query参数进行URL编码后再请求
			finalUrl := fmt.Sprintf("%s://%s%s?%s",
				parsedURL.Scheme,
				parsedURL.Host,
				parsedURL.Path,
				url.PathEscape(parsedURL.RawQuery),
			)
			return finalUrl
		}

		if parsedQuery != nil {
			// 遍历新的查询参数，添加到查询参数中
			for _, param := range r.Params {
				for key, value := range param {
					parsedQuery.Add(key, value)
				}
			}

			// 为URL添加查询参数
			if len(parsedQuery) > 0 {
				finalUrl := strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1),
					parsedQuery.Encode()}, "?")
				return finalUrl
			}
			// 得到最终的URL
			finalUrl := strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)
			return finalUrl
		}
	}
	return userURL
}

func (r *Requests) GetResponse() {
	r.Response.StatusCode = r.HttpResponse.StatusCode                // 响应状态码
	r.Response.EndTime = int(time.Now().UnixNano())                  // 请求结束时间
	r.Response.SpendTime = r.Response.EndTime - r.Response.StartTime // 请求消耗时间（纳秒）
	r.Response.SpendTimeSeconds = r.Response.SpendTime / 1000 / 1000 / 1000

	// 记录请求详情
	if r.Config.IsRecordRequestDetail && r.HttpResponse != nil && r.HttpResponse.Request != nil {
		requestDump, err := httputil.DumpRequest(r.HttpResponse.Request, true)
		if err != nil {
			r.Log.Error("获取请求详情失败", "error", err)
			return
		}
		r.Response.RawReqDetail = string(requestDump)
	}

	// 记录响应详情
	if r.Config.IsRecordResponseDetail && r.HttpResponse != nil {
		responseDump, err := httputil.DumpResponse(r.HttpResponse, true)
		if err != nil {
			r.Log.Error("获取响应详情失败", "error", err)
			return
		}
		r.Response.RawRespDetail = string(responseDump)
	}

	// 获取响应体真实内容
	if r.HttpResponse.Body != nil {
		var Body = r.HttpResponse.Body
		if r.HttpResponse.Header.Get("Content-Encoding") == "gzip" && r.HttpResponse.Header.Get("Accept-Encoding") != "" {
			reader, err := gzip.NewReader(Body)
			if err != nil {
				r.Log.Error("解压响应体内容失败", "error", err)
				return
			}
			Body = reader
		}

		// 读取响应体内容
		content, err := ioutil.ReadAll(Body)
		if err != nil {
			r.Log.Error("读取响应体内容失败", "error", err)
			return
		}

		// 文本内容
		r.Response.Content = content
		r.Response.Text = string(r.Response.Content)
	}
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
