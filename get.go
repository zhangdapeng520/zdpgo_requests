package zdpgo_requests

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func (r *Requests) GetHttpRequest() (req *http.Request) {
	req = &http.Request{
		Method:     "GET",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	// 设置请求头
	req.Header = *r.Header

	// 返回
	return
}

// GetHttpClient 获取HTTP请求的客户端
func (r *Requests) GetHttpClient() (httpClient *http.Client) {
	// 是否跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.CheckHttps},
	}

	// 创建客户端
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(r.Config.Timeout),
	}

	// 超时控制
	if r.Config.Timeout != 0 {
		httpClient.Timeout = time.Duration(r.Config.Timeout) * time.Second
	}

	// 自动生成cookie
	jar, err := cookiejar.New(nil)
	if err != nil {
		r.Log.Error("创建cookie失败", "error", err)
	}
	httpClient.Jar = jar

	// 返回
	return
}

// GetParsedUrl 获取解析后的URL地址
func (r *Requests) GetParsedUrl(userURL string) (finalUrl string, err error) {
	finalUrl = userURL
	var (
		parsedURL   *url.URL
		parsedQuery url.Values
	)

	// 解析URL
	parsedURL, err = url.Parse(userURL)
	if err != nil {
		r.Log.Error("解析URL地址失败", "error", err, "userURL", userURL)
		if !r.Config.IsIgnoredParsedError {
			return
		}
	}

	// 解析Query查询参数
	parsedQuery, err = url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		r.Log.Error("解析query查询参数失败", "error", err, "query", parsedURL.RawQuery)
		if !r.Config.IsIgnoredParsedError { // 不忽略解析错误
			// 无法正常解析query参数，尝试将query参数进行URL编码后再请求
			finalUrl = fmt.Sprintf("%s://%s%s?%s",
				parsedURL.Scheme,
				parsedURL.Host,
				parsedURL.Path,
				url.PathEscape(parsedURL.RawQuery),
			)
		}
		return
	}

	// 遍历新的查询参数，添加到查询参数中
	r.Log.Debug("处理查询参数", "params", r.Params)
	for _, param := range r.Params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}

	// 为URL添加查询参数
	r.Log.Debug("为URL添加查询参数", "parsedQuery", parsedQuery)
	if len(parsedQuery) > 0 {
		finalUrl = strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1),
			parsedQuery.Encode()}, "?")
		r.Log.Debug("获取最终的URL成功", "finalUrl", finalUrl, "parsedURL", parsedURL)
		return
	}

	// 得到最终的URL
	finalUrl = strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)
	r.Log.Debug("获取最终的URL成功", "finalUrl", finalUrl, "parsedURL", parsedURL)
	return
}

func (r *Requests) GetResponse(resp *Response, res *http.Response) *Response {
	resp.StatusCode = res.StatusCode               // 响应状态码
	resp.EndTime = int(time.Now().UnixNano())      // 请求结束时间
	resp.SpendTime = resp.EndTime - resp.StartTime // 请求消耗时间（纳秒）
	resp.SpendTimeSeconds = resp.SpendTime / 1000 / 1000 / 1000

	// 记录请求详情
	if r.Config.IsRecordRequestDetail {
		requestDump, err := httputil.DumpRequest(res.Request, true)
		if err != nil {
			r.Log.Error("获取请求详情失败", "error", err)
			return resp
		}
		resp.RawReqDetail = string(requestDump)
	}

	// 记录响应详情
	if r.Config.IsRecordResponseDetail {
		responseDump, err := httputil.DumpResponse(res, true)
		if err != nil {
			r.Log.Error("获取响应详情失败", "error", err)
			return resp
		}
		resp.RawRespDetail = string(responseDump)
	}

	resp.R = res   // 解析响应
	resp.Content() // 读取内容
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			r.Log.Error("关闭响应体失败", "error", err)
		}
	}(res.Body)
	return resp
}
