package requests

import "net/http"

// SetCookie 设置cookie
func (req *Request) SetCookie(cookie *http.Cookie) {
	req.Cookies = append(req.Cookies, cookie)
}

// ClearCookies 清除cookie
func (req *Request) ClearCookies() {
	req.Cookies = req.Cookies[0:0]
}

// ClientSetCookies 客户端设置cookie
func (req *Request) ClientSetCookies() {
	if len(req.Cookies) > 0 {
		req.Client.Jar.SetCookies(req.httpreq.URL, req.Cookies)
		req.ClearCookies()
	}

}

// Cookies 获取响应的cookie
func (resp *Response) Cookies() (cookies []*http.Cookie) {
	httpreq := resp.req.httpreq
	client := resp.req.Client
	cookies = client.Jar.Cookies(httpreq.URL)
	return cookies
}