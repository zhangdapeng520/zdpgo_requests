package zdpgo_requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

// Any 发送任意请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (req *Request) Any(method, originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	var (
		jsonStr      string // json字符串
		jsonStrBytes []byte // json字符串字节数组
	)

	// 请求的方法
	req.httpReq.Method = strings.ToUpper(method)

	// 清空Header
	for k, _ := range req.httpReq.Header {
		delete(req.httpReq.Header, k)
	}

	// 设置默认的请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var params []map[string]string // query查询参数
	var datas []map[string]string  // form表单数据
	var files []map[string]string  // 文件列表

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) { // 已经自动转换了真实的类型
		case Header: // 设置请求头
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case Params: //	设置Query查询参数
			params = append(params, a)
		case Datas: // 设置POST数据
			datas = append(datas, a)
		case Files: //	设置文件列表
			files = append(files, a)
		case BaseAuth: //	设置权限校验
			req.httpReq.SetBasicAuth(a.Username, a.Password)
		case map[string]string: // 如果是map，默认当data数据处理
			datas = append(datas, a)
		case JsonData: // 如果是JsonData结构体类型
			jsonStrBytes, err = json.Marshal(arg.(JsonData))
			if err != nil {
				Log.Error("解析Json数据失败", "error", err)
				return nil, err
			}
			req.Header.Set("Content-Type", "application/json")
			req.setBodyRawBytes(ioutil.NopCloser(strings.NewReader(string(jsonStrBytes))))
		case JsonString: //	如果是Json字符串
			req.Header.Set("Content-Type", "application/json")
			jsonStr = string(arg.(JsonString))
			req.setBodyRawBytes(ioutil.NopCloser(strings.NewReader(jsonStr)))
		case string: //	如果是字符串，则当成是raw纯文本数据
			req.setBodyRawBytes(ioutil.NopCloser(strings.NewReader(arg.(string))))
		}
	}

	// 构建目标地址
	destUrl, _ := buildURLParams(originUrl, ignoreParseError, params...)

	if len(files) > 0 {
		req.buildFilesAndForms(files, datas) // 构建文件和表单
	} else if len(datas) > 0 {
		Forms := req.buildForms(datas...) // 构建表单
		req.setBodyBytes(Forms)
	}

	// 准备执行请求
	URL, err := url.Parse(destUrl)
	if err != nil {
		Log.Error("解析目标地址失败", "error", err, "destUrl", destUrl)
		return nil, err
	}
	req.httpReq.URL = URL
	req.ClientSetCookies()

	// 构建请求对象
	resp = &Response{
		StartTime: int(time.Now().UnixNano()),
	}
	req.Client.CheckRedirect = func(req1 *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			resp.IsRedirect = true
			resp.RedirectUrl = req1.URL.String()
		}
		return http.ErrUseLastResponse
	}
	res, err := req.Client.Do(req.httpReq)
	if err != nil {
		Log.Error("发送请求失败", "error", err)
		return nil, err
	}
	resp.StatusCode = res.StatusCode               // 响应状态码
	resp.EndTime = int(time.Now().UnixNano())      // 请求结束时间
	resp.SpendTime = resp.EndTime - resp.StartTime // 请求消耗时间（纳秒）
	resp.SpendTimeSeconds = resp.SpendTime / 1000 / 1000 / 1000

	// 记录请求详情
	requestDump, err := httputil.DumpRequest(res.Request, true)
	if err != nil {
		return nil, err
	}
	resp.RawReqDetail = string(requestDump)

	// 记录响应详情
	responseDump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}
	resp.RawRespDetail = string(responseDump)

	// 请求请求内容
	req.httpReq.Body = nil        // 清空请求体
	req.httpReq.GetBody = nil     // 清空get参数
	req.httpReq.ContentLength = 0 // 清空内容长度

	// 解析响应
	resp.R = res
	resp.req = req

	// 读取内容
	resp.Content()
	defer res.Body.Close()

	// 返回响应
	return resp, nil
}

// Any 发送任意请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (r *Requests) Any(method, originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response,
	err error) {
	var (
		jsonStrBytes []byte // json字符串字节数组
	)

	// 请求的方法
	r.HttpReq.Method = strings.ToUpper(method)

	// 清空Header
	for k, _ := range r.HttpReq.Header {
		delete(r.HttpReq.Header, k)
	}

	var params []map[string]string // query查询参数
	var datas []map[string]string  // form表单数据
	var files []map[string]string  // 文件列表

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) { // 已经自动转换了真实的类型
		case Header: // 设置请求头
			for k, v := range a {
				r.Header.Set(k, v)
			}
		case Params: //	设置Query查询参数
			r.Params = append(r.Params, a)
		case Datas: // 设置POST数据
			r.Forms = append(r.Forms, a)
		case Files: //	设置文件列表
			r.Files = append(r.Files, a)
		case BaseAuth: //	设置权限校验
			r.HttpReq.SetBasicAuth(a.Username, a.Password)
		case map[string]string: // 如果是map，默认当data数据处理
			r.Forms = append(r.Forms, a)
		case JsonData: // 如果是JsonData结构体类型
			jsonStrBytes, err = json.Marshal(a)
			if err != nil {
				Log.Error("解析Json数据失败", "error", err)
				return nil, err
			}
			r.Header.Set("Content-Type", "application/json")
			r.SetBodyByBytes(jsonStrBytes)
		case JsonString: //	如果是Json字符串
			r.Header.Set("Content-Type", "application/json")
			r.SetBodyByString(string(a))
		case string: //	如果是字符串，则当成是raw纯文本数据
			r.SetBodyByString(a)
		}
	}

	// 构建目标地址
	destUrl, err := buildURLParams(originUrl, ignoreParseError, params...)
	if err != nil {
		r.Log.Error("构建目标地址失败", "error", err, "originUrl", originUrl)
	}

	// 构建文件和表单
	if len(files) > 0 {
		r.SetFilesAndForms(files, datas) // 设置文件和表单
	} else if len(datas) > 0 {
		r.SetForms(datas...) // 设置表单
	}

	// 准备执行请求
	URL, err := url.Parse(destUrl)
	if err != nil {
		Log.Error("解析目标地址失败", "error", err, "destUrl", destUrl)
		return nil, err
	}
	r.HttpReq.URL = URL
	r.SetCookies()

	// 构建请求对象
	resp = &Response{
		StartTime: int(time.Now().UnixNano()),
	}
	r.Client.CheckRedirect = func(req1 *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			resp.IsRedirect = true
			resp.RedirectUrl = req1.URL.String()
		}
		return http.ErrUseLastResponse
	}

	// 执行请求
	res, err := r.Client.Do(r.HttpReq)
	if err != nil {
		Log.Error("发送请求失败", "error", err)
		return nil, err
	}
	resp.StatusCode = res.StatusCode               // 响应状态码
	resp.EndTime = int(time.Now().UnixNano())      // 请求结束时间
	resp.SpendTime = resp.EndTime - resp.StartTime // 请求消耗时间（纳秒）
	resp.SpendTimeSeconds = resp.SpendTime / 1000 / 1000 / 1000

	// 记录请求详情
	requestDump, err := httputil.DumpRequest(res.Request, true)
	if err != nil {
		return nil, err
	}
	resp.RawReqDetail = string(requestDump)

	// 记录响应详情
	responseDump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}
	resp.RawRespDetail = string(responseDump)

	// 请求请求内容
	r.HttpReq.Body = nil        // 清空请求体
	r.HttpReq.GetBody = nil     // 清空get参数
	r.HttpReq.ContentLength = 0 // 清空内容长度

	// 解析响应
	resp.R = res
	resp.reqs = r

	// 读取内容
	resp.Content()
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			r.Log.Error("关闭响应体失败", "error", err)
		}
	}(res.Body)

	// 返回响应
	return resp, nil
}

// AnyJsonWithTimeout 发送任意请求并携带JSON数据
func (r *Requests) AnyJsonWithTimeout(method string, targetUrl string, body map[string]interface{},
	timeout int) (result string,
	err error) {

	// 准备json字符串
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return
	}

	// 准备请求对象
	req, err := http.NewRequest(strings.ToUpper(method), targetUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "ZDPGo-Requests")

	// 准备客户端
	cli := http.Client{
		Timeout: time.Second * time.Duration(timeout), // 超时时间
	}
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取结果
	b, _ := io.ReadAll(resp.Body)
	result = string(b)
	return
}

// Get 发送GET请求
func (r *Requests) Get(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Get(url, false, args...)
	return resp, err
}

// GetIgnoreParseError 发送GET请求，且忽略解析URL时遇到的错误
func (r *Requests) GetIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Get(url, true, args...)
	return resp, err
}

// GetHttpRequest 获取HTTP请求对象
func (r *Requests) GetHttpRequest1(reqMethod, reqUrl string, requestBody io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(strings.ToUpper(reqMethod), reqUrl, requestBody)
	if err != nil {
		r.Log.Error("创建HTTP请求对象失败", "error", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", r.Config.ContentType)
	req.Header.Set("User-Agent", r.Config.UserAgent)

	// 返回
	return
}

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Post(url, false, args...)
	return resp, err
}

// PostIgnoreParseError 发送POST请求，且忽略解析URL时遇到的错误
func (r *Requests) PostIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Post(url, true, args...)
	return resp, err
}

// PostJsonWithTimeout 发送JSON请求并携带JSON数据
func (r *Requests) PostJsonWithTimeout(targetUrl string, body map[string]interface{}, timeout int) (result string, err error) {
	return r.AnyJsonWithTimeout("post", targetUrl, body, timeout)
}

// Patch 发送PATCH请求
func (r *Requests) Patch(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Patch(url, false, args...)
	return resp, err
}

// PatchIgnoreParseError 发送PATCH请求，且忽略解析URL时遇到的错误
func (r *Requests) PatchIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Patch(url, true, args...)
	return resp, err
}

// Put 发送PUT请求
func (r *Requests) Put(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Put(url, false, args...)
	return resp, err
}

// PutIgnoreParseError 发送PUT请求，且忽略解析URL时遇到的错误
func (r *Requests) PutIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Put(url, true, args...)
	return resp, err
}

// Delete 发送DELETE请求
func (r *Requests) Delete(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Delete(url, false, args...)
	return resp, err
}

// DeleteIgnoreParseError 发送DELETE请求，且忽略解析URL时遇到的错误
func (r *Requests) DeleteIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Request.Delete(url, true, args...)
	return resp, err
}
