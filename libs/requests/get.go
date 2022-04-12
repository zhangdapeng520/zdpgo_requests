package requests

import (
	"fmt"
	"net/url"
)

// Get 发送GET请求
func (req *Request) Get(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	// 设置请求方法为GET
	req.httpreq.Method = "GET"

	// 设置参数 ?a=b&b=c
	var params []map[string]string

	// 重置cookie
	// Client.Do can copy cookie from client.Jar to req.Header
	delete(req.httpreq.Header, "Cookie")

	// 遍历窜进来的参数
	for _, arg := range args {
		// 检测参数的类型
		switch a := arg.(type) {
		// 如果是请求头类型：requests.Header
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
			// arg is "GET" params
			// ?title=website&id=1860&from=login
		//	如果是参数类型：requests.Params
		case Params:
			params = append(params, a)
		//	如果是权限校验类型
		case Auth:
			// a{username,password}
			req.httpreq.SetBasicAuth(a[0], a[1])
		}
	}

	// 构建请求参数
	destUrl, _ := buildURLParams(originUrl, ignoreParseError, params...)

	// 解析目标地址
	URL, err := url.Parse(destUrl)
	if err != nil {
		return nil, err
	}
	req.httpreq.URL = URL

	// 设置cookie
	req.ClientSetCookies()
	req.RequestDebug()

	// 发送请求，获取响应
	res, err := req.Client.Do(req.httpreq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 解析响应
	resp = &Response{}
	resp.R = res
	resp.req = req

	// 读取内容
	resp.Content()
	defer res.Body.Close()

	resp.ResponseDebug()
	return resp, nil
}
