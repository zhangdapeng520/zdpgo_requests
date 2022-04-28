package zdpgo_requests

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

// Proxy 设置代理
func (req *Request) Proxy(proxyUrl string) error {
	// 创建url对象
	urli := url.URL{}

	// 解析代理
	urlProxy, err := urli.Parse(proxyUrl)
	if err != nil {
		return err
	}

	// 设置客户端的代理
	req.Client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(urlProxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return nil
}
