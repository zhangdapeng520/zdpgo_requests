package zdpgo_requests

import (
	"bytes"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadToResponse 上传文件并获取响应
// @param targetUrl 目标地址
// @param formName 文件表单名称
// @param filePath 上传文件的路径
// @return 响应对象，错误对象
func (r *Requests) UploadToResponse(targetUrl string, formName string, filePath string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile(formName, filePath)
	if err != nil {
		return nil, err
	}

	// 打开文件句柄操作
	fh, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Upload 上传文件
// @param targetUrl 目标地址
// @param filePath 上传文件的路径
// @return 错误对象
func (r *Requests) Upload(targetUrl string, filePath string) error {
	resp, err := r.UploadToResponse(targetUrl, "file", filePath)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// UploadToBytes 上传文件并返回响应体字节数组
// @param targetUrl 目标地址
// @param filePath 上传文件的路径
// @return 响应内容，错误对象
func (r *Requests) UploadToBytes(targetUrl string, filePath string) ([]byte, error) {
	resp, err := r.UploadToResponse(targetUrl, "file", filePath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体数据
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// UploadFsToBytes 上传FS文件系统的文件，返回bytes数据
func (r *Requests) UploadFsToBytes(targetUrl string, fsObj fs.FS, fileFormName, filePath string) (result []byte,
	err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile(fileFormName, filePath)
	if err != nil {
		return
	}

	// 打开文件句柄操作
	fh, err := fsObj.Open(filePath)
	if err != nil {
		return
	}
	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return
	}

	// 读取响应体数据
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 返回结果
	return
}

// UploadFsToString 上传FS系统文件，并将结果转换为字符串返回
func (r *Requests) UploadFsToString(targetUrl string, fsObj fs.FS, fileFormName, filePath string) (result string,
	err error) {
	toBytes, err := r.UploadFsToBytes(targetUrl, fsObj, fileFormName, filePath)
	if err != nil {
		return
	}
	result = string(toBytes)
	return
}

// UploadToString 上传文件并返回响应字符串
// @param targetUrl 目标地址
// @param filePath 上传文件的路径
// @return 响应内容，错误对象
func (r *Requests) UploadToString(targetUrl string, filePath string) (string, error) {
	resp, err := r.UploadToBytes(targetUrl, filePath)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
