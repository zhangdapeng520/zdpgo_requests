package zdpgo_requests

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

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
	req.httpReq.Body = ioutil.NopCloser(bytes.NewReader(b.Bytes()))
	req.httpReq.ContentLength = int64(b.Len())
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
			// 无法正常解析query参数，尝试讲query参数进行URL编码后再请求
			resultUrl := fmt.Sprintf("%s://%s%s?%s",
				parsedURL.Scheme,
				parsedURL.Host,
				parsedURL.Path,
				url.PathEscape(parsedURL.RawQuery),
			)
			return resultUrl, nil
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
