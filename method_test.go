package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_log"

var (
	proxyUrl = "http://127.0.0.1:8080"
	r        = getRequests()
)

func getRequests() *Requests {
	r := NewWithConfig(&Config{
		Timeout:  5,
		ProxyUrl: proxyUrl,
	}, zdpgo_log.Tmp)
	return r
}
