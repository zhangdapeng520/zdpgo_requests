package zdpgo_requests

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// Response 响应对象
type Response struct {
	R                *http.Response // 响应对象
	content          []byte         // 响应内容
	text             string         // 响应文本
	req              *Request       // 请求对象
	RawReqDetail     string         // 请求详情字符串
	RawRespDetail    string         // 响应详情字符串
	StatusCode       int            // 状态码
	IsRedirect       bool           // 是否重定向了
	RedirectUrl      string         // 重定向的的URL地址
	StartTime        int            // 请求开始时间（纳秒）
	EndTime          int            // 请求结束时间（纳秒）
	SpendTime        int            // 请求消耗时间（纳秒）
	SpendTimeSeconds int            // 请求消耗时间（秒）
}

// Content 获取响应内容
func (resp *Response) Content() []byte {
	var err error

	if len(resp.content) > 0 {
		return resp.content
	}

	var Body = resp.R.Body
	if resp.R.Header.Get("Content-Encoding") == "gzip" && resp.req.Header.Get("Accept-Encoding") != "" {
		reader, err := gzip.NewReader(Body)
		if err != nil {
			return nil
		}
		Body = reader
	}

	resp.content, err = ioutil.ReadAll(Body)
	if err != nil {
		return nil
	}

	return resp.content
}

// Text 获取响应文本
func (resp *Response) Text() string {
	// 没有响应
	if resp == nil {
		return ""
	}

	// 获取响应
	if resp.content == nil {
		resp.Content()
	}

	// 获取响应文本
	resp.text = string(resp.content)
	return resp.text
}

// SaveFile 保存文件
func (resp *Response) SaveFile(filename string) error {
	if resp.content == nil {
		resp.Content()
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(resp.content)
	f.Sync()

	return err
}

// Json 解析响应内容为json数据
func (resp *Response) Json(v interface{}) error {
	if resp.content == nil {
		resp.Content()
	}
	return json.Unmarshal(resp.content, v)
}
