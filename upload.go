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
func (r *Requests) Upload(urlPath, formName, filePath string) (*Response, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		r.Log.Error("打开文件失败", "error", err)
		return nil, err
	}
	response, err := r.UploadByBytes(urlPath, formName, filePath, fileContent)
	if err != nil {
		r.Log.Error("上传文件失败", "error", err, "urlPath", urlPath, "filePath", filePath)
		return nil, err
	}
	return response, nil
}

// UploadByBytes 上传字节数组
func (r *Requests) UploadByBytes(urlPath, formName, fileName string, fileContent []byte) (*Response, error) {
	response := &Response{}

	// 创建缓冲器
	bodyBuffer := &bytes.Buffer{}

	// 创建写入对象
	bodyWriter := multipart.NewWriter(bodyBuffer)

	// 创建文件
	fileWriter, err := bodyWriter.CreateFormFile(formName, fileName)
	if err != nil {
		r.Log.Error("创建表单失败", "error", err)
		return nil, err
	}

	// 赋值文件到写入对象
	io.Copy(fileWriter, strings.NewReader(string(fileContent)))

	// 获取请求头
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// 构建请求对象
	client := r.GetHttpClient()
	request := &Request{
		Method: "POST",
		Url:    urlPath,
		Header: map[string]string{
			"Content-Type": contentType,
		},
		Body: bodyBuffer,
	}
	r.setHeader(request)
	req := r.GetHttpRequest(*request)

	// 设置请求体内容
	req.ContentLength = int64(bodyBuffer.Len())
	buf := bodyBuffer.Bytes()
	req.GetBody = func() (io.ReadCloser, error) {
		r := bytes.NewReader(buf)
		return io.NopCloser(r), nil
	}
	bodyReader := ioutil.NopCloser(bodyBuffer)
	req.Body = bodyReader

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		r.Log.Error("上传文件失败", "error", err)
		return nil, err
	}

	// 设置响应结果
	r.SetResponse(response, resp)

	return response, nil
}
