package zdpgo_requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Any 发送任意请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (r *Requests) Any(method, originUrl string, args ...interface{}) (*Response, error) {
	response := &Response{}

	r.InitData() // 初始化数据

	defer func() {
		// 删除临时目录
		if r.File.IsExists(r.Config.TmpDir) {
			r.File.DeleteDir(r.Config.TmpDir)
		}

		// 捕获异常
		if err := recover(); err != nil {
			r.Log.Error("处理请求失败", "error", err)
		}
	}()

	// 请求的方法
	r.HttpReq.Method = strings.ToUpper(method)

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) { // 已经自动转换了真实的类型
		case Header: // 设置请求头
			for k, v := range a {
				r.Header.Set(k, v)
			}
		case BaseAuth: // 设置权限校验
			r.HttpReq.SetBasicAuth(a.Username, a.Password)
		case http.Cookie: // 设置cookie
			r.Cookies = append(r.Cookies, &a)
		case Param: //	添加查询参数
			r.Params = append(r.Params, a)
		case Files: // 添加上传文件
			r.Files = append(r.Files, a)
		case FormFileBytes: // 添加上传文件
			r.FileBytesList = append(r.FileBytesList, a)
		case map[string]string: // 如果是map，默认当data数据处理
			r.Forms = append(r.Forms, a)
		case JsonMap: // 如果是JsonData结构体类型
			jsonStrBytes, err := json.Marshal(a)
			if err != nil {
				r.Log.Error("解析Json数据失败", "error", err)
				return nil, err
			}
			r.Header.Set("Content-Type", "application/json")
			r.SetBodyByBytes(jsonStrBytes)
		case JsonString: //	如果是Json字符串
			r.Header.Set("Content-Type", "application/json")
			r.SetBodyByString(string(a))
		case string: //	如果是字符串，则当成是raw纯文本数据
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.SetBodyByString(a)
		}
	}

	// 构建目标地址
	destUrl := r.GetParsedUrl(originUrl)
	r.SetFilesAndForms() // 构建文件和表单

	// 准备执行请求
	URL, err := url.Parse(destUrl)
	if err != nil {
		r.Log.Error("解析目标地址失败", "error", err, "destUrl", destUrl)
		return nil, err
	}
	r.HttpReq.URL = URL
	r.SetCookies()
	r.HttpReq.Header = *r.Header

	// 构建请求对象
	r.Response.StartTime = int(time.Now().UnixNano())
	r.Client.CheckRedirect = func(req1 *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			r.Response.IsRedirect = true
			r.Response.RedirectUrl = req1.URL.String()
		}
		return http.ErrUseLastResponse
	}
	if r.Config.IsKeepSession {
		r.ClientSetCookies()
	}

	// 执行请求
	r.HttpResponse, err = r.Client.Do(r.HttpReq)
	if err != nil {
		r.Log.Error("发送请求失败", "error", err)
		return nil, err
	}

	// 获取响应信息
	r.SetResponse(response, r.HttpResponse)

	// 返回响应
	return response, nil
}

// 不解析URL发送请求
func (r *Requests) AnyNoParseURL(request Request) (*Response, error) {
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

// Get 发送GET请求
func (r *Requests) Get(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("get", url, args...)
	return resp, err
}

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("post", url, args...)
	return resp, err
}

// Patch 发送PATCH请求
func (r *Requests) Patch(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("patch", url, args...)
	return resp, err
}

// Put 发送PUT请求
func (r *Requests) Put(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("put", url, args...)
	return resp, err
}

// Delete 发送DELETE请求
func (r *Requests) Delete(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("delete", url, args...)
	return resp, err
}
