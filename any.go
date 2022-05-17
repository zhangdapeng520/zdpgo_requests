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
