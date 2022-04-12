package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// Post 发送POST请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func (req *Request) Post(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	// 请求的方法
	req.httpreq.Method = "POST"

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

	disturl, _ := buildURLParams(originUrl, ignoreParseError, params...)

	if len(files) > 0 {
		// 构建文件和表单
		req.buildFilesAndForms(files, datas)
	} else {
		// 构建表单
		Forms := req.buildForms(datas...)
		req.setBodyBytes(Forms) // set forms to body
	}

	// 准备执行请求
	URL, err := url.Parse(disturl)
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

	resp = &Response{}
	resp.R = res
	resp.req = req

	resp.Content()
	defer res.Body.Close()

	return resp, nil
}

// Post 发送POST请求
// @param url 要请求的URL路径
// @param ignoreParseError 是否忽略解析URL错误
// @param args 要携带的参数
func Post(url string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// 调用req.POST处理请求
	resp, err = req.Post(url, ignoreParseError, args...)
	return resp, err
}

// PostJson 发送POST请求且传递json格式的数据
func PostJson(originUrl string, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// 发送req的POST请求
	resp, err = req.PostJson(originUrl, args...)
	return resp, err
}

// PostJson 发送POST请求
func (req *Request) PostJson(origurl string, args ...interface{}) (resp *Response, err error) {
	// 设置请求方法
	req.httpreq.Method = "POST"

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 重置cookie
	// Client.Do can copy cookie from client.Jar to req.Header
	delete(req.httpreq.Header, "Cookie")

	// 遍历参数
	for _, arg := range args {
		switch a := arg.(type) {
		// 设置请求头
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		//	设置数据内容，期望是一个json字符串
		case string:
			req.setBodyRawBytes(ioutil.NopCloser(strings.NewReader(arg.(string))))
		//	设置权限校验
		case Auth:
			// a{username,password}
			req.httpreq.SetBasicAuth(a[0], a[1])
		//	其他数据类型，一律使用json数据传递
		default:
			b := new(bytes.Buffer)
			err = json.NewEncoder(b).Encode(a)
			if err != nil {
				return nil, err
			}
			req.setBodyRawBytes(ioutil.NopCloser(b))
		}
	}

	//prepare to Do
	URL, err := url.Parse(origurl)
	if err != nil {
		return nil, err
	}
	req.httpreq.URL = URL

	req.ClientSetCookies()

	res, err := req.Client.Do(req.httpreq)

	// clear post  request information
	req.httpreq.Body = nil
	req.httpreq.GetBody = nil
	req.httpreq.ContentLength = 0

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp = &Response{}
	resp.R = res
	resp.req = req

	resp.Content()
	defer res.Body.Close()
	return resp, nil
}
