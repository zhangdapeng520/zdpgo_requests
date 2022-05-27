package zdpgo_requests

import (
	"testing"
)

var (
	urlPath  = "http://localhost:3333/ping"
	proxyUrl = "http://10.1.3.12:8080"
	jsonUrl  = "http://localhost:3333/json"
	formUrl  = "http://localhost:3333/form"
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
	var data []Request
	data = append(data, Request{
		Method: "GET",
		Url:    urlPath,
	})
	data = append(data, Request{
		Method: "POST",
		Url:    urlPath,
	})

	data = append(data, Request{
		Method: "DELETE",
		Url:    urlPath,
	})
	data = append(data, Request{
		Method: "PUT",
		Url:    urlPath,
	})
	data = append(data, Request{
		Method: "PATCH",
		Url:    urlPath,
	})

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
	jsonText := "{\"text\":\"Hello\"}"
	jsonHeader := map[string]string{"Content-Type": "application/json"}
	var data []Request
	data = append(data, Request{
		Method: "POST",
		Url:    jsonUrl,
		Json:   jsonData,
	})

	data = append(data, Request{
		Method: "DELETE",
		Url:    jsonUrl,
		Json:   jsonData,
	})
	data = append(data, Request{
		Method: "PUT",
		Url:    jsonUrl,
		Json:   jsonData,
	})
	data = append(data, Request{
		Method: "PATCH",
		Url:    jsonUrl,
		Json:   jsonData,
	})

	data = append(data, Request{
		Method:   "POST",
		Header:   jsonHeader,
		Url:      jsonUrl,
		JsonText: jsonText,
	})

	data = append(data, Request{
		Method:   "DELETE",
		Header:   jsonHeader,
		Url:      jsonUrl,
		JsonText: jsonText,
	})
	data = append(data, Request{
		Method:   "PUT",
		Header:   jsonHeader,
		Url:      jsonUrl,
		JsonText: jsonText,
	})
	data = append(data, Request{
		Method:   "PATCH",
		Header:   jsonHeader,
		Url:      jsonUrl,
		JsonText: jsonText,
	})

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
func TestRequests_AnyForm(t *testing.T) {
	formData := map[string]string{"a": "1", "b": "2.2", "c": "33", "d": "true"}
	formText := "a=b&b=1&b=2.22"
	var data []Request
	data = append(data, Request{
		Method: "GET",
		Url:    formUrl,
		Form:   formData,
	})

	data = append(data, Request{
		Method: "POST",
		Url:    formUrl,
		Form:   formData,
	})

	data = append(data, Request{
		Method: "DELETE",
		Url:    formUrl,
		Form:   formData,
	})
	data = append(data, Request{
		Method: "PUT",
		Url:    formUrl,
		Form:   formData,
	})
	data = append(data, Request{
		Method: "PATCH",
		Url:    formUrl,
		Form:   formData,
	})
	data = append(data, Request{
		Method:   "GET",
		Url:      formUrl,
		FormText: formText,
	})
	data = append(data, Request{
		Method:   "POST",
		Url:      formUrl,
		FormText: formText,
	})

	data = append(data, Request{
		Method:   "DELETE",
		Url:      formUrl,
		FormText: formText,
	})
	data = append(data, Request{
		Method:   "PUT",
		Url:      formUrl,
		FormText: formText,
	})
	data = append(data, Request{
		Method:   "PATCH",
		Url:      formUrl,
		FormText: formText,
	})

	for _, request := range data {
		response, err := r.AnyForm(request)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("状态码不是200")
		}
	}
}

func TestRequests_AnyText(t *testing.T) {
	jsonText := "{\"text\":\"Hello\"}"
	jsonHeader := map[string]string{"Content-Type": "application/json"}
	var data []Request
	data = append(data, Request{
		Method:   "POST",
		Url:      urlPath,
		JsonText: jsonText,
	})

	data = append(data, Request{
		Method:   "DELETE",
		Url:      urlPath,
		JsonText: jsonText,
	})
	data = append(data, Request{
		Method:   "PUT",
		Url:      urlPath,
		JsonText: jsonText,
	})
	data = append(data, Request{
		Method:   "PATCH",
		Url:      urlPath,
		JsonText: jsonText,
	})

	data = append(data, Request{
		Method:   "POST",
		Header:   jsonHeader,
		Url:      urlPath,
		JsonText: jsonText,
	})

	data = append(data, Request{
		Method:   "DELETE",
		Header:   jsonHeader,
		Url:      urlPath,
		JsonText: jsonText,
	})
	data = append(data, Request{
		Method:   "PUT",
		Header:   jsonHeader,
		Url:      urlPath,
		JsonText: jsonText,
	})
	data = append(data, Request{
		Method:   "PATCH",
		Header:   jsonHeader,
		Url:      urlPath,
		JsonText: jsonText,
	})

	for _, request := range data {
		response, err := r.AnyText(request)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("状态码不是200")
		}
	}
}
