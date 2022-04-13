package requests

// Get 发送GET请求
func (req *Request) Get(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	resp, err = req.Any("GET", originUrl, ignoreParseError, args...)
	return resp, err
}

// Get 发送GET请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func Get(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// 调用request发送GET请求
	resp, err = req.Get(originUrl, ignoreParseError, args...)
	return resp, err
}
