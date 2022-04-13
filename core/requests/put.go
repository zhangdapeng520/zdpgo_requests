package requests

// Put 发送PUT请求
func (req *Request) Put(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	resp, err = req.Any("PUT", originUrl, ignoreParseError, args...)
	return resp, err
}

// Put 发送PUT请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func Put(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()
	resp, err = req.Put(originUrl, ignoreParseError, args...)
	return resp, err
}
