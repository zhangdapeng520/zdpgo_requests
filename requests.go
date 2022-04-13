package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_requests/core/requests"

type Requests struct {
	Request *requests.Request // 请求对象
}

func New() *Requests {
	r := Requests{}

	// 实例化请求对象
	r.Request = requests.Requests()

	return &r
}
