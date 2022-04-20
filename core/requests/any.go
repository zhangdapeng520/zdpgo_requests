package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Any 发送任意请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (req *Request) Any(method, originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	// 请求的方法
	req.httpreq.Method = strings.ToUpper(method)

	// 清空Header
	for k, _ := range req.httpreq.Header {
		delete(req.httpreq.Header, k)
	}

	// 设置默认的请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var params []map[string]string // query查询参数
	var datas []map[string]string  // form表单数据
	var files []map[string]string  // 文件列表

	// 重置cookie
	// Client.Do can copy cookie from client.Jar to req.Header
	//delete(req.httpreq.Header, "Cookie")

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) {
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
		case Auth: //	设置权限校验
			// a{username,password}
			req.httpreq.SetBasicAuth(a[0], a[1])
		case map[string]string: // 如果是map，默认当data数据处理
			datas = append(datas, a)
		case JsonData: // 如果是JsonData结构体类型
			jsonStr, err := json.Marshal(arg.(JsonData))
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", "application/json")
			req.setBodyRawBytes(ioutil.NopCloser(strings.NewReader(string(jsonStr))))
		case JsonString: //	如果是Json字符串
			req.Header.Set("Content-Type", "application/json")
			jsonStr := string(arg.(JsonString))
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
		return nil, err
	}
	req.httpreq.URL = URL
	req.ClientSetCookies()

	// 发送请求
	res, err := req.Client.Do(req.httpreq)
	resp = &Response{
		StatusCode: res.StatusCode,
	}

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
	req.httpreq.Body = nil        // 清空请求体
	req.httpreq.GetBody = nil     // 清空get参数
	req.httpreq.ContentLength = 0 // 清空内容长度

	// 解析响应
	resp.R = res
	resp.req = req

	// 读取内容
	resp.Content()
	defer res.Body.Close()

	// 返回响应
	return resp, nil
}
