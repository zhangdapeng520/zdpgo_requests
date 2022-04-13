package requests

// Post 发送POST请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (req *Request) Post(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	resp, err = req.Any("POST", originUrl, ignoreParseError, args...)
	return resp, err
}

// Post 发送POST请求
// @param url 要请求的URL路径
// @param ignoreParseError 是否忽略解析URL错误
// @param args 要携带的参数
func Post(url string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()
	resp, err = req.Post(url, ignoreParseError, args...)
	return resp, err
}
