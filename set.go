package zdpgo_requests

import (
	"compress/gzip"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

/*
@Time : 2022/5/19 15:12
@Author : 张大鹏
@File : set.go
@Software: Goland2021.3.1
@Description: set设置数据相关
*/

// SetProxy 设置代理
func (r *Requests) SetProxy(client *http.Client, proxyUrl string) error {
	// 解析代理地址
	uri, _ := url.Parse(proxyUrl)

	// 设置代理
	client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(uri),                                      // 设置代理
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.IsCheckHttps}, // 是否跳过证书校验
	}
	client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间
	r.Config.ProxyUrl = proxyUrl

	return nil
}

// SetTimeout 设置请求超时时间
func (r *Requests) SetTimeout(client *http.Client, timeout int) {
	if timeout <= 0 {
		timeout = 60
	}
	client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间
}

// SetResponse 设置响应结果
func (r *Requests) SetResponse(resp *Response, response *http.Response) {
	if response == nil {
		return
	}

	if resp == nil {
		resp = &Response{}
	}
	if response == nil {
		return
	}
	resp.StatusCode = response.StatusCode          // 响应状态码
	resp.EndTime = int(time.Now().UnixNano())      // 请求结束时间
	resp.SpendTime = resp.EndTime - resp.StartTime // 请求消耗时间（纳秒）
	resp.SpendTimeSeconds = resp.SpendTime / 1000 / 1000 / 1000

	// 源端口
	resp.ClientPort = r.ClientPort

	// 记录请求详情
	if r.Config.IsRecordRequestDetail && response.Request != nil {
		requestDump, err := httputil.DumpRequest(response.Request, true)
		if err != nil {
			return
		}
		resp.RawReqDetail = string(requestDump)
	}

	// 记录响应详情
	if r.Config.IsRecordResponseDetail && response != nil {
		responseDump, err := httputil.DumpResponse(response, true)
		if err != nil {
			return
		}
		resp.RawRespDetail = string(responseDump)
	}

	// 响应体没有内容
	if response.Body == nil {
		return
	}
	defer response.Body.Close()

	// 获取响应体真实内容
	var Body = response.Body
	if response.Header.Get("Content-Encoding") == "gzip" && response.Header.Get("Accept-Encoding") != "" {
		reader, err := gzip.NewReader(Body)
		if err != nil {
			return
		}
		Body = reader
	}

	// 读取响应体内容
	content, err := ioutil.ReadAll(Body)
	if err != nil {
		return
	}

	// 文本内容
	resp.Content = content
	resp.Text = string(resp.Content)

	// 任务数量减少
	r.TaskNum--
}
