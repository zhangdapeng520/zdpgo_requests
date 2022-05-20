package zdpgo_requests

import (
	"net/http"
)

// SetCookie 设置cookie
func (req *Requests) SetCookie(cookie *http.Cookie) {
	req.Cookies = append(req.Cookies, cookie)
}

// ClearCookies 清除cookie
func (req *Requests) ClearCookies() {
	req.Cookies = req.Cookies[0:0]
}

// ClientSetCookies 客户端设置cookie
func (req *Requests) ClientSetCookies() {
	if len(req.Cookies) > 0 {
		req.Client.Jar.SetCookies(req.HttpReq.URL, req.Cookies)
		req.ClearCookies()
	}

}

// GetCookies 获取响应的cookie
func (r *Requests) GetCookies() (cookies []*http.Cookie) {
	cookies = r.Client.Jar.Cookies(r.HttpReq.URL)
	return cookies
}
