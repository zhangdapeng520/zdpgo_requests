package zdpgo_requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

// Any 发送任意请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (r *Requests) Any(method, originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response,
	err error) {
	defer func() {
		// 删除临时目录
		if r.Exists(r.Config.FsTmpDir) {
			r.DeleteDir(r.Config.FsTmpDir)
		}
	}()

	var (
		jsonStrBytes []byte // json字符串字节数组
	)

	// 请求的方法
	r.HttpReq.Method = strings.ToUpper(method)

	// 清空Header
	for k, _ := range r.HttpReq.Header {
		delete(r.HttpReq.Header, k)
	}

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) { // 已经自动转换了真实的类型
		case Header: // 设置请求头
			for k, v := range a {
				r.Header.Set(k, v)
			}
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
			r.HttpReq.Header.Set("Content-Type", "application/json")
			r.SetBodyByBytes(jsonStrBytes)
		case JsonString: //	如果是Json字符串
			r.HttpReq.Header.Set("Content-Type", "application/json")
			r.SetBodyByString(string(a))
		case string: //	如果是字符串，则当成是raw纯文本数据
			r.SetBodyByString(a)
		}
	}

	// 构建目标地址
	destUrl, err := buildURLParams(originUrl, ignoreParseError, r.Params...)
	if err != nil {
		r.Log.Error("构建目标地址失败", "error", err, "originUrl", originUrl)
	}

	r.SetFilesAndForms() // 构建文件和表单

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
	resp.R = res // 解析响应

	// 读取内容
	resp.Content()
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			r.Log.Error("关闭响应体失败", "error", err)
		}
	}(res.Body)

	r.InitData() // 初始化数据

	// 返回响应
	return resp, nil
}

// AnyJsonWithTimeout 发送任意请求并携带JSON数据
func (r *Requests) AnyJsonWithTimeout(method string, targetUrl string, body map[string]interface{},
	timeout int) (result string,
	err error) {

	defer func() {
		// 删除临时目录
		if r.Exists(r.Config.FsTmpDir) {
			r.DeleteDir(r.Config.FsTmpDir)
		}
	}()

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
	resp, err := r.Any("get", url, false, args...)
	return resp, err
}

// GetIgnoreParseError 发送GET请求，且忽略解析URL时遇到的错误
func (r *Requests) GetIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("get", url, true, args...)
	return resp, err
}

// Post 发送POST请求
func (r *Requests) Post(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("post", url, false, args...)
	return resp, err
}

// PostIgnoreParseError 发送POST请求，且忽略解析URL时遇到的错误
func (r *Requests) PostIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("post", url, true, args...)
	return resp, err
}

// PostJsonWithTimeout 发送JSON请求并携带JSON数据
func (r *Requests) PostJsonWithTimeout(targetUrl string, body map[string]interface{}, timeout int) (result string, err error) {
	return r.AnyJsonWithTimeout("post", targetUrl, body, timeout)
}

// Patch 发送PATCH请求
func (r *Requests) Patch(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("patch", url, false, args...)
	return resp, err
}

// PatchIgnoreParseError 发送PATCH请求，且忽略解析URL时遇到的错误
func (r *Requests) PatchIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("patch", url, true, args...)
	return resp, err
}

// Put 发送PUT请求
func (r *Requests) Put(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("put", url, false, args...)
	return resp, err
}

// PutIgnoreParseError 发送PUT请求，且忽略解析URL时遇到的错误
func (r *Requests) PutIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("put", url, true, args...)
	return resp, err
}

// Delete 发送DELETE请求
func (r *Requests) Delete(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("delete", url, false, args...)
	return resp, err
}

// DeleteIgnoreParseError 发送DELETE请求，且忽略解析URL时遇到的错误
func (r *Requests) DeleteIgnoreParseError(url string, args ...interface{}) (*Response, error) {
	resp, err := r.Any("delete", url, true, args...)
	return resp, err
}

// 打开文件用于上传
func openFile(filename string) *os.File {
	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return r
}

// 处理URL的参数
func buildURLParams(userURL string, ignoreParseError bool, params ...map[string]string) (string, error) {
	// 解析URL
	parsedURL, err := url.Parse(userURL)
	if err != nil {
		return "", err
	}

	// 解析Query查询参数
	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		if ignoreParseError {
			// 无法正常解析query参数，尝试讲query参数进行URL编码后再请求
			resultUrl := fmt.Sprintf("%s://%s%s?%s",
				parsedURL.Scheme,
				parsedURL.Host,
				parsedURL.Path,
				url.PathEscape(parsedURL.RawQuery),
			)
			return resultUrl, nil
		}
		return "", nil
	}

	// 遍历参数，添加到查询参数中
	for _, param := range params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}

	// 为URL添加查询参数
	return addQueryParams(parsedURL, parsedQuery), nil
}

// 为URL添加查询参数
func addQueryParams(parsedURL *url.URL, parsedQuery url.Values) string {
	if len(parsedQuery) > 0 {
		return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?")
	}
	return strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)
}
