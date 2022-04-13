package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/core/requests"

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Post(url, false, args...)
	return resp, err
}

// PostIgnoreParseError 发送POST请求，且忽略解析URL时遇到的错误
func (r *Requests) PostIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Post(url, true, args...)
	return resp, err
}