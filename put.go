package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/core/requests"

// Put 发送PUT请求
func (r *Requests) Put(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Put(url, false, args...)
	return resp, err
}

// PutIgnoreParseError 发送PUT请求，且忽略解析URL时遇到的错误
func (r *Requests) PutIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Put(url, true, args...)
	return resp, err
}
