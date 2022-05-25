package zdpgo_requests

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
@Time : 2022/5/19 15:12
@Author : 张大鹏
@File : set.go
@Software: Goland2021.3.1
@Description: set设置数据相关
*/

// SetBodyByForm 设置请求体为表单字段
func (r *Requests) SetBodyByForm(Forms url.Values) {
	data := Forms.Encode()
	r.HttpReq.Body = ioutil.NopCloser(strings.NewReader(data))
	r.HttpReq.ContentLength = int64(len(data))
}

// SetBodyByReader 设置请求体为二进制输入流
func (r *Requests) SetBodyByReader(read io.ReadCloser) {
	r.HttpReq.Body = read
}

// SetBodyByString 设置字符串为请求体内容
func (r *Requests) SetBodyByString(data string) {
	r.SetBodyByReader(ioutil.NopCloser(strings.NewReader(data)))
}

// SetBodyByBytes 设置字节数组为请求体内容
func (r *Requests) SetBodyByBytes(data []byte) {
	r.SetBodyByString(string(data))
}

// SetFilesAndForms 设置上传文件和表单
func (r *Requests) SetFilesAndForms() {
	if len(r.Files) == 0 && len(r.Forms) == 0 {
		return
	}

	var buffer bytes.Buffer // 处理文件
	tmpDir := r.Config.TmpDir
	writer := multipart.NewWriter(&buffer) // 创建表单对象

	// 字节类型的文件
	if len(r.FileBytesList) > 0 {
		for i, file := range r.FileBytesList {
			if file.FormName == "" {
				if i > 0 {
					file.FormName = fmt.Sprintf("file%d", i)
				} else {
					file.FormName = "file"
				}
			}
			if file.FormName == "" {
				file.FileName = r.Random.Str(32)
			}

			// 先存到临时目录
			tmpFileName := fmt.Sprintf("%s/%s", tmpDir, file.FileName)
			err := ioutil.WriteFile(tmpFileName, file.ContentBytes, 0644)
			if err != nil {
				r.Log.Error("保存临时文件失败", "error", err, "fileName", file.FileName)
				continue
			}
			r.Files = append(r.Files, map[string]string{
				file.FileName: tmpFileName,
			})
		}
	}

	// 遍历文件列表
	if len(r.Files) > 0 {

		// 如果使用了嵌入文件系统，需要将文件先转移到临时目录
		if r.IsFs {
			// 不存在则创建文件
			if !r.File.IsExists(tmpDir) {
				r.File.CreateMultiDir(tmpDir)
			}

			// 保存文件到临时目录
			for _, file := range r.Files {
				for _, v := range file {
					fileName := filepath.Base(v)
					fh, err1 := r.Fs.ReadFile(v)
					if err1 != nil {
						r.Log.Error("读取文件失败", "error", err1, "fileName", v)
						continue
					}
					err := ioutil.WriteFile(fmt.Sprintf("%s/%s", tmpDir, fileName), fh, 0644)
					if err != nil {
						r.Log.Error("保存临时文件失败", "error", err, "fileName", v)
						continue
					}
				}
			}
		}
		for _, file := range r.Files {
			for k, v := range file {
				// 如果使用了FS嵌入文件系统，从临时目录读
				if r.IsFs {
					fileName := filepath.Base(v)
					v = fmt.Sprintf("%s/%s", tmpDir, fileName)
				}

				// 创建文件表单
				part, err := writer.CreateFormFile(k, v)
				if err != nil {
					r.Log.Error("处理要上传的文件失败", "error", err)
				}

				// 复制文件到请求体
				fileObj, err := os.Open(v)
				if err != nil {
					r.Log.Error("打开文件失败", "error", err, "file", v)
				}
				_, err = io.Copy(part, fileObj)
				if err != nil {
					r.Log.Error("复制文件到上次对象失败", "error", err)
				}
				err = fileObj.Close()
				if err != nil {
					r.Log.Error("关闭文件对象失败", "error", err)
				}
			}
		}
	}

	// 添加表单数据
	if len(r.Forms) > 0 {
		for _, data := range r.Forms {
			for k, v := range data {
				err := writer.WriteField(k, v)
				if err != nil {
					r.Log.Error("添加表单数据失败", "error", err, "key", k, "value", v)
				}
			}
		}
	}

	err := writer.Close()
	if err != nil {
		r.Log.Error("关闭表单对象失败", "error", err)
		return
	}

	// 设置文件头："Content-Type": "multipart/form-data; boundary=------------------------7d87eceb5520850c",
	r.HttpReq.Body = ioutil.NopCloser(bytes.NewReader(buffer.Bytes()))
	r.HttpReq.ContentLength = int64(buffer.Len())
	r.Header.Set("Content-Type", writer.FormDataContentType())
}

// SetForms 设置表单数据
func (r *Requests) SetForms(datas ...map[string]string) {
	form := url.Values{}
	for _, data := range datas {
		for key, value := range data {
			form.Add(key, value)
		}
	}
	r.SetBodyByForm(form)
}

// SetCookies 客户端设置cookie
func (r *Requests) SetCookies() {
	if len(r.Cookies) > 0 {
		// 客户端设置cookie
		r.Client.Jar.SetCookies(r.HttpReq.URL, r.Cookies)

		// 清除请求对象的cookie
		r.Cookies = r.Cookies[0:0]
	}
}

// SetProxy 设置代理
func (r *Requests) SetProxy(proxyUrl string) error {
	// 解析代理地址
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		r.Log.Error("解析代理地址失败", "error", err, "proxyUrl", proxyUrl)
	}

	// 设置代理
	r.Client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(uri),                                    // 设置代理
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.CheckHttps}, // 是否跳过证书校验
	}
	r.Client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间

	return nil
}

// SetTimeout 设置请求超时时间
func (r *Requests) SetTimeout(timeout int) {
	if timeout <= 0 {
		timeout = 60
	}
	r.Client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间
}

// SetResponse 设置响应结果
func (r *Requests) SetResponse(resp *Response, response *http.Response) {
	if resp == nil {
		resp = &Response{}
	}
	if response == nil {
		r.Log.Warning("HTTP响应为空，无法处理", "response", response)
		return
	}
	resp.StatusCode = response.StatusCode                      // 响应状态码
	resp.EndTime = int(time.Now().UnixNano())                  // 请求结束时间
	resp.SpendTime = r.Response.EndTime - r.Response.StartTime // 请求消耗时间（纳秒）
	resp.SpendTimeSeconds = r.Response.SpendTime / 1000 / 1000 / 1000

	// 源端口
	resp.ClientPort = r.ClientPort

	// 记录请求详情
	if r.Config.IsRecordRequestDetail && r.HttpResponse != nil && r.HttpResponse.Request != nil {
		requestDump, err := httputil.DumpRequest(r.HttpResponse.Request, true)
		if err != nil {
			r.Log.Error("获取请求详情失败", "error", err)
			return
		}
		resp.RawReqDetail = string(requestDump)
	}

	// 记录响应详情
	if r.Config.IsRecordResponseDetail && r.HttpResponse != nil {
		responseDump, err := httputil.DumpResponse(r.HttpResponse, true)
		if err != nil {
			r.Log.Error("获取响应详情失败", "error", err)
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
			r.Log.Error("解压响应体内容失败", "error", err)
			return
		}
		Body = reader
	}

	// 读取响应体内容
	content, err := ioutil.ReadAll(Body)
	if err != nil {
		r.Log.Error("读取响应体内容失败", "error", err)
		return
	}

	// 文本内容
	resp.Content = content
	resp.Text = string(resp.Content)

	// 将响应结果挂载到对象上
	r.Response = resp
	r.HttpResponse = nil // 置空
}
