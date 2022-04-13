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
	resp, err = req.Any("POST", originUrl, ignoreParseError, args...)
	return resp, err
}

// Post 发送POST请求
// @param url 要请求的URL路径
// @param ignoreParseError 是否忽略解析URL错误
// @param args 要携带的参数
func Post(url string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()
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
