package zdpgo_requests

import (
	"encoding/json"
	"net/http"
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

// setHeader 设置请求头
func (r *Requests) setHeader(request *Request) {
	if request.Header == nil {
		if r.Config.IsRandomUserAgent {
			request.Header = map[string]string{
				"User-Agent": r.GetRandomUserAgent(),
			}
		} else {
			request.Header = map[string]string{
				"User-Agent": r.Config.UserAgent,
			}
		}
		if request.IsText {
			request.Header["Content-Type"] = "text/plain"
		} else if request.IsForm {
			request.Header["Content-Type"] = "application/x-www-form-urlencoded"
		} else if request.IsJson {
			request.Header["Content-Type"] = "application/json"
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
			if request.IsText {
				request.Header["Content-Type"] = "text/plain"
			} else if request.IsForm {
				request.Header["Content-Type"] = "application/x-www-form-urlencoded"
			} else if request.IsJson {
				request.Header["Content-Type"] = "application/json"
			} else {
				request.Header["Content-Type"] = r.Config.ContentType
			}
		}
	}

	// 自定义请求头
	request.Header["X-Author"] = r.Config.Author
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

	// 设置请求头
	r.setHeader(request)

	// 响应对象
	response := &Response{}

	// http请求对象
	if request.Method == "" {
		request.Method = "GET"
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

// AnyCompareStatusCode 任意方法发送请求，会发送两次请求，比较前后的状态码
func (r *Requests) AnyCompareStatusCode(method, target1Url, target2Url string, args ...interface{}) (*Response, error) {
	// 发送第一次请求
	response1, err := r.Any(method, target1Url)
	if err != nil {
		r.Log.Error("发送第一次请求失败", "error", err)
		return response1, err
	}

	// 发送第二次请求
	response2, err := r.Any(method, target2Url, args...)
	if err != nil {
		r.Log.Error("发送第二次请求失败", "error", err)
		return response2, err
	}

	// 更新状态码
	response2.FirstStatusCode = response1.StatusCode
	return response2, nil
}

// Get 发送GET请求
func (r *Requests) Get(targetUrl string, args ...interface{}) (*Response, error) {
	return r.Any("GET", targetUrl, args...)
}

// Post 发送POST请求
func (r *Requests) Post(targetUrl string, args ...interface{}) (*Response, error) {
	return r.Any("POST", targetUrl, args...)
}

// Put 发送PUT请求
func (r *Requests) Put(targetUrl string, args ...interface{}) (*Response, error) {
	return r.Any("PUT", targetUrl, args...)
}

// Patch 发送PATCH请求
func (r *Requests) Patch(targetUrl string, args ...interface{}) (*Response, error) {
	return r.Any("PATCH", targetUrl, args...)
}

// Delete 发送DELETE请求
func (r *Requests) Delete(targetUrl string, args ...interface{}) (*Response, error) {
	return r.Any("DELETE", targetUrl, args...)
}
