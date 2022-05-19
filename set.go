package zdpgo_requests

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"strings"
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
// @param files 要上传的文件列表
// @param datas 要上传的表单数据
func (r *Requests) SetFilesAndForms(files []map[string]string, datas []map[string]string) {
	// 处理文件
	var b bytes.Buffer

	// 创建表单对象
	w := multipart.NewWriter(&b)

	// 遍历文件列表
	for _, file := range files {
		for k, v := range file {
			// 创建文件表单
			part, err := w.CreateFormFile(k, v)
			if err != nil {
				r.Log.Error("处理要上传的文件失败", "error", err)
			}

			// 复制文件到请求体
			fileObj := openFile(v)
			_, err = io.Copy(part, fileObj)
			if err != nil {
				r.Log.Error("复制文件到上次对象失败", "error", err)
			}
		}
	}

	// 添加表单数据
	for _, data := range datas {
		for k, v := range data {
			err := w.WriteField(k, v)
			if err != nil {
				r.Log.Error("添加表单数据失败", "error", err, "key", k, "value", v)
			}
		}
	}
	err := w.Close()
	if err != nil {
		r.Log.Error("关闭表单对象失败", "error", err)
		return
	}

	// 设置文件头："Content-Type": "multipart/form-data; boundary=------------------------7d87eceb5520850c",
	r.HttpReq.Body = ioutil.NopCloser(bytes.NewReader(b.Bytes()))
	r.HttpReq.ContentLength = int64(b.Len())
	r.Header.Set("Content-Type", w.FormDataContentType())
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
