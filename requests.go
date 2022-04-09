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
	resp, err := requests.Get(url, args...)
	return resp, err
}

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*requests.Response, error) {
	resp, err := requests.Post(url, args...)
	return resp, err
}
