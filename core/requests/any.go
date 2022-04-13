package requests

import (
	"io/ioutil"
	"net/url"
	"strings"
)

// Any 发送任意请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (req *Request) Any(method, originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	// 请求的方法
	req.httpreq.Method = strings.ToUpper(method)

	// 设置默认的请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set params ?a=b&b=c
	var params []map[string]string // query查询参数
	var datas []map[string]string  // form表单数据
	var files []map[string]string  // 文件列表
	var rawData string             // 纯文本数据

	// 重置cookie
	// Client.Do can copy cookie from client.Jar to req.Header
	delete(req.httpreq.Header, "Cookie")

	// 遍历请求参数
	for _, arg := range args {
		switch a := arg.(type) {
		// 设置请求头
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
			// ?title=website&id=1860&from=login
		//	设置Query查询参数
		case Params:
			params = append(params, a)
		// 设置POST数据
		case Datas: //Post form data,packaged in body.
			datas = append(datas, a)
		//	设置文件列表
		case Files:
			files = append(files, a)
		//	设置权限校验
		case Auth:
			// a{username,password}
			req.httpreq.SetBasicAuth(a[0], a[1])
		//	如果是map，默认当data数据处理
		case map[string]string:
			datas = append(datas, a)
		//	如果是字符串，则当成是raw纯文本数据
		case string:
			rawData = a
		}
	}

	// 构建目标地址
	destUrl, _ := buildURLParams(originUrl, ignoreParseError, params...)

	if len(files) > 0 {
		// 构建文件和表单
		req.buildFilesAndForms(files, datas)
	} else {
		// 构建表单
		Forms := req.buildForms(datas...)
		req.setBodyBytes(Forms) // set forms to body
	}

	// 准备执行请求
	URL, err := url.Parse(destUrl)
	if err != nil {
		return nil, err
	}
	req.httpreq.URL = URL
	req.ClientSetCookies()

	// 如果存在纯文本数据，则设置纯文本数据
	if rawData != "" {
		req.setBodyRawBytes(ioutil.NopCloser(strings.NewReader(rawData)))
	}

	// 发送请求
	res, err := req.Client.Do(req.httpreq)

	// 清空POST的参数
	req.httpreq.Body = nil
	req.httpreq.GetBody = nil
	req.httpreq.ContentLength = 0

	if err != nil {
		return nil, err
	}

	// 解析响应
	resp = &Response{}
	resp.R = res
	resp.req = req

	// 读取内容
	resp.Content()
	defer res.Body.Close()

	return resp, nil
}
