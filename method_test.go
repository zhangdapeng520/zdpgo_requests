package zdpgo_requests

import (
	"testing"
)

var (
	urlPath  = "http://localhost:3333/ping"
	proxyUrl = "http://10.1.3.12:8080"
	jsonUrl  = "http://localhost:3333/json"
	r        = getRequests()
)

func getRequests() *Requests {
	r := NewWithConfig(Config{
		Debug:    true,
		Timeout:  5,
		ProxyUrl: proxyUrl,
	})
	return r
}

// 任意类型的方法，不解析URL路径
func TestRequests_Any(t *testing.T) {
	data := []Request{
		{"GET", urlPath, nil, nil, nil, nil, BasicAuth{}},
		{"GET", urlPath, nil, nil, map[string]string{"a": "b"}, nil, BasicAuth{}},
		{"POST", urlPath, nil, nil, nil, nil, BasicAuth{}},
		{"DELETE", urlPath, nil, nil, nil, nil, BasicAuth{}},
		{"PUT", urlPath, nil, nil, nil, nil, BasicAuth{}},
		{"PATCH", urlPath, nil, nil, nil, nil, BasicAuth{}},
	}

	for _, request := range data {
		response, err := r.Any(request)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("状态码不是200")
		}
	}
}

func TestRequests_AnyJson(t *testing.T) {
	jsonData := map[string]interface{}{"a": 1, "b": 2.2, "c": "33", "d": true}
	data := []Request{
		{"POST", jsonUrl, nil, nil, nil, jsonData, BasicAuth{}},
		{"DELETE", jsonUrl, nil, nil, nil, jsonData, BasicAuth{}},
		{"PUT", jsonUrl, nil, nil, nil, jsonData, BasicAuth{}},
		{"PATCH", jsonUrl, nil, nil, nil, jsonData, BasicAuth{}},
	}

	for _, request := range data {
		response, err := r.AnyJson(request)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("状态码不是200")
		}
	}
}
