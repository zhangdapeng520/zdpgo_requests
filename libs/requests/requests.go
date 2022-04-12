/* Copyright（2） 2018 by  asmcos .
Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package requests

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

var VERSION string = "0.8"

type Request struct {
	httpreq *http.Request
	Header  *http.Header
	Client  *http.Client
	Debug   int
	Cookies []*http.Cookie
}

type Response struct {
	R       *http.Response
	content []byte
	text    string
	req     *Request
}

type Header map[string]string
type Params map[string]string
type Datas map[string]string // POST提交的数据
type Files map[string]string // 文件列表：name ,filename

// Auth 权限校验类型，{username,password}
type Auth []string

func Requests() *Request {

	req := new(Request)

	req.httpreq = &http.Request{
		Method:     "GET",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	req.Header = &req.httpreq.Header
	req.httpreq.Header.Set("User-Agent", "Go-Requests "+VERSION)

	req.Client = &http.Client{}

	// auto with Cookies
	// cookiejar.New source code return jar, nil
	jar, _ := cookiejar.New(nil)

	req.Client.Jar = jar

	return req
}

// Get 发送GET请求
// @param originUrl 要请求的URL地址
// @param args 请求携带的参数
func Get(originUrl string, ignoreParseError bool, args ...interface{}) (resp *Response, err error) {
	req := Requests()

	// 调用request发送GET请求
	resp, err = req.Get(originUrl, ignoreParseError, args...)
	return resp, err
}

// 处理URL的参数
func buildURLParams(userURL string, ignoreParseError bool, params ...map[string]string) (string, error) {
	// 解析URL
	parsedURL, err := url.Parse(userURL)
	if err != nil {
		return "", err
	}

	// 解析Query查询参数
	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		if ignoreParseError {
			return userURL, nil
		}
		return "", nil
	}

	// 遍历参数，添加到查询参数中
	for _, param := range params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}

	// 为URL添加查询参数
	return addQueryParams(parsedURL, parsedQuery), nil
}

// 为URL添加查询参数
func addQueryParams(parsedURL *url.URL, parsedQuery url.Values) string {
	if len(parsedQuery) > 0 {
		return strings.Join([]string{strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1), parsedQuery.Encode()}, "?")
	}
	return strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)
}

// RequestDebug 请求调试
func (req *Request) RequestDebug() {
	if req.Debug != 1 {
		return
	}
	fmt.Println("===========Go RequestDebug ============")
	message, err := httputil.DumpRequestOut(req.httpreq, false)
	if err != nil {
		return
	}
	fmt.Println(string(message))

	if len(req.Client.Jar.Cookies(req.httpreq.URL)) > 0 {
		fmt.Println("Cookies:")
		for _, cookie := range req.Client.Jar.Cookies(req.httpreq.URL) {
			fmt.Println(cookie)
		}
	}
}

// cookies
// cookies only save to Client.Jar
// req.Cookies is temporary
func (req *Request) SetCookie(cookie *http.Cookie) {
	req.Cookies = append(req.Cookies, cookie)
}

func (req *Request) ClearCookies() {
	req.Cookies = req.Cookies[0:0]
}

func (req *Request) ClientSetCookies() {

	if len(req.Cookies) > 0 {
		// 1. Cookies have content, Copy Cookies to Client.jar
		// 2. Clear  Cookies
		req.Client.Jar.SetCookies(req.httpreq.URL, req.Cookies)
		req.ClearCookies()
	}

}

// set timeout s = second
func (req *Request) SetTimeout(n time.Duration) {
	req.Client.Timeout = time.Duration(n * time.Second)
}

func (req *Request) Close() {
	req.httpreq.Close = true
}

func (req *Request) Proxy(proxyurl string) {

	urli := url.URL{}
	urlproxy, err := urli.Parse(proxyurl)
	if err != nil {
		fmt.Println("Set proxy failed")
		return
	}
	req.Client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(urlproxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

}

/**************/
func (resp *Response) ResponseDebug() {

	if resp.req.Debug != 1 {
		return
	}

	fmt.Println("===========Go ResponseDebug ============")

	message, err := httputil.DumpResponse(resp.R, false)
	if err != nil {
		return
	}

	fmt.Println(string(message))

}

func (resp *Response) Content() []byte {

	var err error

	if len(resp.content) > 0 {
		return resp.content
	}

	var Body = resp.R.Body
	if resp.R.Header.Get("Content-Encoding") == "gzip" && resp.req.Header.Get("Accept-Encoding") != "" {
		// fmt.Println("gzip")
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

func (resp *Response) Text() string {
	if resp.content == nil {
		resp.Content()
	}
	resp.text = string(resp.content)
	return resp.text
}

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

func (resp *Response) Json(v interface{}) error {
	if resp.content == nil {
		resp.Content()
	}
	return json.Unmarshal(resp.content, v)
}

func (resp *Response) Cookies() (cookies []*http.Cookie) {
	httpreq := resp.req.httpreq
	client := resp.req.Client

	cookies = client.Jar.Cookies(httpreq.URL)

	return cookies

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

	req.RequestDebug()

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
	resp.ResponseDebug()
	return resp, nil
}

// 设置表单的字段
func (req *Request) setBodyBytes(Forms url.Values) {
	data := Forms.Encode()
	req.httpreq.Body = ioutil.NopCloser(strings.NewReader(data))
	req.httpreq.ContentLength = int64(len(data))
}

// 设置表单的二进制输入流
func (req *Request) setBodyRawBytes(read io.ReadCloser) {
	req.httpreq.Body = read
}

// 构建文件和表单
func (req *Request) buildFilesAndForms(files []map[string]string, datas []map[string]string) {
	// 处理文件
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// 变量文件列表
	for _, file := range files {
		for k, v := range file {
			part, err := w.CreateFormFile(k, v)
			if err != nil {
				fmt.Printf("上传文件 %s 失败！", v)
				panic(err)
			}
			file := openFile(v)
			_, err = io.Copy(part, file)
			if err != nil {
				panic(err)
			}
		}
	}

	// 添加表单数据
	for _, data := range datas {
		for k, v := range data {
			w.WriteField(k, v)
		}
	}
	w.Close()

	// 设置文件头："Content-Type": "multipart/form-data; boundary=------------------------7d87eceb5520850c",
	req.httpreq.Body = ioutil.NopCloser(bytes.NewReader(b.Bytes()))
	req.httpreq.ContentLength = int64(b.Len())
	req.Header.Set("Content-Type", w.FormDataContentType())
}

// 构建表单数据
func (req *Request) buildForms(datas ...map[string]string) (Forms url.Values) {
	Forms = url.Values{}
	for _, data := range datas {
		for key, value := range data {
			Forms.Add(key, value)
		}
	}
	return Forms
}

// 打开文件用于上传
func openFile(filename string) *os.File {
	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return r
}
