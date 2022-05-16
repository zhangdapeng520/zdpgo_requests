package zdpgo_requests

import (
	"io"
	"net/http"
	"strings"
)

// Get 发送GET请求
func (r *Requests) Get(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Get(url, false, args...)
	return resp, err
}

// GetIgnoreParseError 发送GET请求，且忽略解析URL时遇到的错误
func (r *Requests) GetIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Get(url, true, args...)
	return resp, err
}

// GetHttpRequest 获取HTTP请求对象
func (r *Requests) GetHttpRequest(reqMethod, reqUrl string, requestBody io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(strings.ToUpper(reqMethod), reqUrl, requestBody)
	if err != nil {
		r.Log.Error("创建HTTP请求对象失败", "error", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", r.Config.ContentType)
	req.Header.Set("User-Agent", r.Config.UserAgent)

	// 返回
	return
}

// GetHttpClient 获取HTTP请求的客户端
func (r *Requests) GetHttpClient() *http.Client {
	var httpClient = &http.Client{}
	return httpClient
}
