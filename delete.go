package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/core/requests"

// Delete 发送DELETE请求
func (r *Requests) Delete(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Delete(url, false, args...)
	return resp, err
}

// DeleteIgnoreParseError 发送DELETE请求，且忽略解析URL时遇到的错误
func (r *Requests) DeleteIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Delete(url, true, args...)
	return resp, err
}