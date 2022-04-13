package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/core/requests"

// Get 发送GET请求
func (r *Requests) Get(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Get(url, false, args...)
	return resp, err
}

// GetIgnoreParseError 发送GET请求，且忽略解析URL时遇到的错误
func (r *Requests) GetIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Get(url, true, args...)
	return resp, err
}
