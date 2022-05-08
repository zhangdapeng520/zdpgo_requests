package zdpgo_requests

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Post(url, false, args...)
	return resp, err
}

// PostIgnoreParseError 发送POST请求，且忽略解析URL时遇到的错误
func (r *Requests) PostIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Post(url, true, args...)
	return resp, err
}

// PostJsonWithTimeout 发送JSON请求并携带JSON数据
func (r *Requests) PostJsonWithTimeout(targetUrl string, body map[string]interface{}, timeout int) (result string, err error) {
	return r.AnyJsonWithTimeout("post", targetUrl, body, timeout)
}
