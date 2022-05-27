package zdpgo_requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 不解析URL发送请求
func (r *Requests) Any(request Request) (*Response, error) {
	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

	// 校验目标地址
	if request.Url == "" {
		return nil, errors.New("目标URL地址不能为空")
	}

	// 响应对象
	response := &Response{}

	// http请求对象
	if request.Method == "" {
		request.Method = "GET"
	}
	if request.Header == nil {
		request.Header = map[string]string{
			"User-Agent":   r.Config.UserAgent,
			"Content-Type": r.Config.ContentType,
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header["User-Agent"] = r.Config.UserAgent
		}
		if _, ok := request.Header["Content-Type"]; !ok {
			request.Header["Content-Type"] = r.Config.ContentType
		}
	}

	req := r.GetHttpRequest(request)

	// 构建请求对象
	response.StartTime = int(time.Now().UnixNano())

	// 获取客户端对象
	client := r.GetHttpClient()
	if r.Config.IsCheckRedirect {
		client.CheckRedirect = func(req1 *http.Request, via []*http.Request) error {
			if len(via) > 0 {
				response.IsRedirect = true
				response.RedirectUrl = req1.URL.String()
			}
			return http.ErrUseLastResponse
		}
	}

	// 执行请求
	httpResponse, err := client.Do(req)
	if err != nil {
		r.Log.Error("发送请求失败", "error", err)
		return nil, err
	}

	// 获取响应信息
	r.SetResponse(response, httpResponse)

	// 返回响应
	return response, nil
}

// 不解析URL发送请求，携带JSON数据
func (r *Requests) AnyJson(request Request) (*Response, error) {
	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

	// 校验json
	if request.Json == nil {
		return nil, errors.New("JSON数据不能为空")
	}

	// 校验方法
	if request.Method == "" {
		request.Method = "POST"
	}

	// 校验目标地址
	if request.Url == "" {
		return nil, errors.New("目标URL地址不能为空")
	}

	// 响应对象
	response := &Response{}

	// http请求对象
	if request.Method == "" {
		request.Method = "GET"
	}
	if request.Header == nil {
		request.Header = map[string]string{
			"User-Agent":   r.Config.UserAgent,
			"Content-Type": "application/json",
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header["User-Agent"] = r.Config.UserAgent
		}
		request.Header["Content-Type"] = "application/json"
	}

	req := r.GetHttpRequest(request)

	// 处理json数据
	dataByte, err := json.Marshal(request.Json)
	if err != nil {
		r.Log.Error("解析JSON数据失败", "error", err, "data", request.Json)
		return nil, err
	}
	bodyReader := ioutil.NopCloser(strings.NewReader(string(dataByte)))
	req.Body = bodyReader

	// 构建请求对象
	response.StartTime = int(time.Now().UnixNano())

	// 获取客户端对象
	client := r.GetHttpClient()
	client.CheckRedirect = func(req1 *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			response.IsRedirect = true
			response.RedirectUrl = req1.URL.String()
		}
		return http.ErrUseLastResponse
	}

	// 执行请求
	httpResponse, err := client.Do(req)
	if err != nil {
		r.Log.Error("发送请求失败", "error", err)
		return nil, err
	}

	// 获取响应信息
	r.SetResponse(response, httpResponse)

	// 返回响应
	return response, nil
}
