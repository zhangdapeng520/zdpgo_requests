package zdpgo_requests

// Delete 发送DELETE请求
func (req *Request) Delete(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	resp, err = req.Any("DELETE", originUrl, ignoreParseError, args...)
	return resp, err
}

// Delete 发送DELETE请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func Delete(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := NewRequest()
	resp, err = req.Delete(originUrl, ignoreParseError, args...)
	return resp, err
}
