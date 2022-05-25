package zdpgo_requests

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

/*
@Time : 2022/5/25 17:33
@Author : 张大鹏
@File : upload.go
@Software: Goland2021.3.1
@Description:
*/

// Upload 普通文件上传
func (r *Requests) Upload(urlPath, formName, filePath string) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		r.Log.Error("打开文件失败", "error", err)
		return
	}
	r.UploadByBytes(urlPath, formName, filePath, fileContent)
}

// UploadByBytes 上传字节数组
func (r *Requests) UploadByBytes(urlPath, formName, fileName string, fileContent []byte) {
	// 创建缓冲器
	bodyBuffer := &bytes.Buffer{}

	// 创建写入对象
	bodyWriter := multipart.NewWriter(bodyBuffer)

	// 创建文件
	fileWriter, err := bodyWriter.CreateFormFile(formName, fileName)
	if err != nil {
		r.Log.Error("创建表单失败", "error", err)
		return
	}

	// 赋值文件到写入对象
	io.Copy(fileWriter, strings.NewReader(string(fileContent)))

	// 获取请求头
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// 构建请求对象
	client := r.GetHttpClient()
	req := r.GetHttpRequest(Request{
		Method: "POST",
		Url:    urlPath,
		Header: map[string]string{
			"Content-Type": contentType,
			"User-Agent":   r.Config.UserAgent,
		},
		Body: bodyBuffer,
	})

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		r.Log.Error("上传文件失败", "error", err)
		return
	}

	// 设置响应结果
	r.SetResponse(&Response{}, resp)
}
