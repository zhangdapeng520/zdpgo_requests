package zdpgo_requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) { // 已经自动转换了真实的类型
		case Header: // 设置请求头
			for k, v := range a {
				r.Header.Set(k, v)
			}
		case BaseAuth: //	设置权限校验
			r.HttpReq.SetBasicAuth(a.Username, a.Password)
		case Param: //	添加查询参数
			r.Params = append(r.Params, a)
		case Files: // 添加上传文件
			r.Files = append(r.Files, a)
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
	destUrl, err := r.GetParsedUrl(originUrl)
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

	resp = r.GetResponse(resp, res) // 获取响应信息
	r.InitData()                    // 初始化数据

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
