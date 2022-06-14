package zdpgo_requests

var (
	proxyUrl = "http://127.0.0.1:8080"
	r        = getRequests()
)

func getRequests() *Requests {
	r := NewWithConfig(&Config{
		Debug:    true,
		Timeout:  5,
		ProxyUrl: proxyUrl,
	})
	return r
}
