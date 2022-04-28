package zdpgo_requests

type Requests struct {
	Request *Request // 请求对象
}

func New() *Requests {
	r := Requests{}

	// 实例化请求对象
	r.Request = NewRequest()

	return &r
}
