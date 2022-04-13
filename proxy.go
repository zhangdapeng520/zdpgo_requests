package zdpgo_requests

// SetProxy 设置代理
func (r *Requests) SetProxy(proxyUrl string) error {
	err := r.Request.Proxy(proxyUrl)
	return err
}
