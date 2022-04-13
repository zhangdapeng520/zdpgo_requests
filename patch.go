package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/core/requests"

// Patch 发送PATCH请求
func (r *Requests) Patch(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Patch(url, false, args...)
	return resp, err
}

// PatchIgnoreParseError 发送PATCH请求，且忽略解析URL时遇到的错误
func (r *Requests) PatchIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := r.Request.Patch(url, true, args...)
	return resp, err
}
