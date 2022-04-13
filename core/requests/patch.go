package requests

// Patch 发送PATCH请求
func (req *Request) Patch(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	resp, err = req.Any("PATCH", originUrl, ignoreParseError, args...)
	return resp, err
}

// Patch 发送PATCH请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func Patch(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()
	resp, err = req.Patch(originUrl, ignoreParseError, args...)
	return resp, err
}
