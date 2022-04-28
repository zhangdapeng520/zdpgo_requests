package zdpgo_requests

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

// Request 请求对象
type Request struct {
	httpreq *http.Request  // http请求对象
	Header  *http.Header   // 请求头
	Client  *http.Client   // 请求客户端
	Debug   int            // 是否为DEBUG模式
	Cookies []*http.Cookie // cookie
	Config  *Config        // 配置对象
}

// NewRequest 创建请求对象
func NewRequest() *Request {
	return NewRequestWithConfig(Config{
		Timeout:       30,
		CheckRedirect: true,
	})
}
func NewRequestWithConfig(config Config) *Request {
	req := new(Request)
	req.httpreq = &http.Request{
		Method:     "GET",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	// 设置请求头
	req.Header = &req.httpreq.Header
	req.httpreq.Header.Set("User-Agent", "ZDPGo-Requests "+VERSION)

	// 设置客户端
	req.Client = &http.Client{}
	req.Config = &config
	if config.Timeout != 0 {
		req.Client.Timeout = time.Duration(config.Timeout) * time.Second
	}

	// 自动生成cookie
	// cookiejar.New source code return jar, nil
	jar, _ := cookiejar.New(nil)
	req.Client.Jar = jar

	// 返回请求对象
	return req
}

// SetTimeout 设置客户端超时时间（秒）
func (req *Request) SetTimeout(n time.Duration) {
	req.Client.Timeout = time.Duration(n * time.Second)
}

// Close 关闭连接
func (req *Request) Close() {
	req.httpreq.Close = true
}
