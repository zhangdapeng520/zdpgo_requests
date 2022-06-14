package zdpgo_requests

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// parseArgs 解析自定义的参数
func (r *Requests) parseArgs(request *Request, args ...interface{}) {
	for _, arg := range args {
		switch argValue := arg.(type) {
		case Header: // 请求头
			if request.Header == nil {
				request.Header = make(map[string]string)
			}
			for k, v := range argValue {
				request.Header[k] = v
			}
		case Text: // 纯文本
			request.IsText = true
			request.Text = string(argValue)
		case string: // 内容
			request.Text = argValue
		case Form: // 表单数据
			request.IsForm = true
			if request.Form == nil {
				request.Form = make(map[string]string)
			}
			for k, v := range argValue {
				request.Form[k] = v
			}
		case map[string]interface{}: // json数据
			request.IsJson = true
			if request.Json == nil {
				request.Json = make(map[string]interface{})
			}
			for k, v := range argValue {
				request.Json[k] = v
			}
		default: // 结构体类型，作为JSON数据处理
			request.IsJson = true
			jsonBytes, err := json.Marshal(argValue)
			if err != nil {
				r.Log.Error("解析JSON数据失败", "error", err)
			} else {
				request.JsonText = string(jsonBytes)
			}
		}
	}
}

// Any 任意方法的请求
func (r *Requests) Any(method, targetUrl string, args ...interface{}) (*Response, error) {
	request := &Request{
		Method: strings.ToUpper(method),
		Url:    targetUrl,
	}
	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

	// 处理参数
	r.parseArgs(request, args...)

	// 响应对象
	response := &Response{}

	// http请求对象
	if request.Method == "" {
		request.Method = "GET"
	}
	if request.Header == nil {
		request.Header = map[string]string{
			"User-Agent": r.Config.UserAgent,
		}
		if request.IsText {
			request.Header["Content-Type"] = "text/plain"
		} else if request.IsForm {
			request.Header["Content-Type"] = "application/x-www-form-urlencoded"
		} else {
			request.Header["Content-Type"] = r.Config.ContentType
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			if r.Config.IsRandomUserAgent {
				request.Header["User-Agent"] = r.GetRandomUserAgent()
			} else {
				request.Header["User-Agent"] = r.Config.UserAgent
			}
		}
		if _, ok := request.Header["Content-Type"]; !ok {
			if r.Config.IsText {
				request.Header["Content-Type"] = "text/plain"
			} else if r.Config.IsForm {
				request.Header["Content-Type"] = "application/x-www-form-urlencoded"
			} else {
				request.Header["Content-Type"] = r.Config.ContentType
			}
		}
	}
	req := r.GetHttpRequest(*request)

	// 构建响应对象
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
		return response, err
	}

	// 获取响应信息
	r.SetResponse(response, httpResponse)

	// 返回响应
	return response, nil
}

// AnyJson 任意方法发送JSON请求
func (r *Requests) AnyJson(request Request) (*Response, error) {
	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

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
	if request.Header == nil {
		request.Header = map[string]string{
			"User-Agent":   r.Config.UserAgent,
			"Content-Type": "application/json",
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header["User-Agent"] = r.Config.UserAgent
		}

	}
	if r.Config.IsText {
		request.Header["Content-Type"] = "application/json"
	} else if r.Config.IsForm {
		request.Header["Content-Type"] = "application/json"
	} else {
		request.Header["Content-Type"] = "application/json"
	}

	req := r.GetHttpRequest(request)

	// 处理json数据
	if request.Json != nil {
		dataByte, err := json.Marshal(request.Json)
		if err != nil {
			r.Log.Error("解析JSON数据失败", "error", err, "data", request.Json)
			return nil, err
		}
		bodyReader := ioutil.NopCloser(strings.NewReader(string(dataByte)))
		req.Body = bodyReader
	} else if request.JsonText != "" {
		bodyReader := ioutil.NopCloser(strings.NewReader(request.JsonText))
		req.Body = bodyReader
	} else {
		return nil, errors.New("JSON数据不能为空")
	}

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

// AnyForm 任意方法发送表单请求
func (r *Requests) AnyForm(request Request) (*Response, error) {
	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

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
	if request.Header == nil {
		request.Header = map[string]string{
			"User-Agent":   r.Config.UserAgent,
			"Content-Type": "application/x-www-form-urlencoded",
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header["User-Agent"] = r.Config.UserAgent
		}
		request.Header["Content-Type"] = "application/x-www-form-urlencoded"
	}

	req := r.GetHttpRequest(request)

	// 处理json数据
	if request.Form != nil {
		data := make(url.Values)
		for key, value := range request.Form {
			data[key] = []string{value}
		}
		bodyReader := strings.NewReader(data.Encode())
		req.ContentLength = int64(bodyReader.Len())
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bodyReader), nil
		}
		req.Body = io.NopCloser(bodyReader)
	} else if request.FormText != "" {
		bodyReader := strings.NewReader(request.FormText)
		req.ContentLength = int64(bodyReader.Len())
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bodyReader), nil
		}
		req.Body = io.NopCloser(bodyReader)
	} else {
		return nil, errors.New("Form表单数据不能为空")
	}

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

// AnyText 任意方法发送纯文本数据
func (r *Requests) AnyText(request Request) (*Response, error) {
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
			"Content-Type": "text/plain",
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header["User-Agent"] = r.Config.UserAgent
		}
		if _, ok := request.Header["Content-Type"]; !ok {
			request.Header["Content-Type"] = "text/plain"
		}
	}
	req := r.GetHttpRequest(request)

	// 处理文本数据
	bodyReader := ioutil.NopCloser(strings.NewReader(request.Text))
	req.Body = bodyReader

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

// Get 发送GET请求
func (r *Requests) Get(targetUrl string) (*Response, error) {
	return r.Any("GET", targetUrl)
}

func (r *Requests) Post(targetUrl string) (*Response, error) {
	return r.Any("POST", targetUrl)
}

var (
	FirstTargetCode = make(map[string]int)
)

// AnyCompareStatus 发送任意方法的文本请求，且必然有Response,会发两次请求，比较状态码
func (r *Requests) AnyCompareStatus(method, firstTarget, secondTarget string) (*Response, error) {
	request := &Request{}

	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

	r.Get(firstTarget)

	// 响应对象
	response := &Response{}

	// http请求对象
	if request.Method == "" {
		request.Method = "GET"
	}
	if request.Header == nil {
		request.Header = map[string]string{
			"Content-Type": "text/plain",
			"X-Author":     r.Config.Author,
		}
		if r.Config.IsRandomUserAgent {
			request.Header["User-Agent"] = r.GetRandomUserAgent()
		} else {
			request.Header["User-Agent"] = r.Config.UserAgent
		}
	} else {
		if _, ok := request.Header["User-Agent"]; !ok {
			request.Header["User-Agent"] = r.Config.UserAgent
		}
		if _, ok := request.Header["Content-Type"]; !ok {
			request.Header["Content-Type"] = "text/plain"
		}
	}

	// 第一个请求对象
	request.Url = firstTarget
	req := r.GetHttpRequest(*request)

	// 处理文本数据
	bodyReader := ioutil.NopCloser(strings.NewReader(request.Text))
	req.Body = bodyReader

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
	response.ClientPort = r.ClientPort

	// 执行第一次请求

	// 执行第二次请求
	httpResponse, err := client.Do(req)
	if err != nil {
		r.Log.Error("发送请求失败", "error", err)
		return response, err
	}

	// 获取响应信息
	r.SetResponse(response, httpResponse)

	// 返回响应
	return response, nil
}
