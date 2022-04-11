package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/libs/requests"

type Requests struct {
}

func New() *Requests {
	r := Requests{}

	return &r
}

// Get 发送GET请求
func (r *Requests) Get(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Get(url, false, args...)
	return resp, err
}

// GetIgnoreParseError 发送GET请求，且忽略解析URL时遇到的错误
func (r *Requests) GetIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Get(url, true, args...)
	return resp, err
}

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Post(url, false, args...)
	return resp, err
}

// PostIgnoreParseError 发送POST请求，且忽略解析URL时遇到的错误
func (r *Requests) PostIgnoreParseError(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Post(url, true, args...)
	return resp, err
}
