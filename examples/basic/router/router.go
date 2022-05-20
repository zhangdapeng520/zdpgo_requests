package router

import (
	"embed"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

/*
@Time : 2022/5/20 9:50
@Author : 张大鹏
@File : router.go
@Software: Goland2021.3.1
@Description: 测试路由相关的用法
*/

func Any(r *zdpgo_requests.Requests, url string) {
	resp, err := r.Any("get", url, true)
	println(resp.Text, err)
	resp, err = r.Any("post", url, true)
	println(resp.Text, err)
	resp, err = r.Any("put", url, true)
	println(resp.Text, err)
	resp, err = r.Any("delete", url, true)
	println(resp.Text, err)
	resp, err = r.Any("patch", url, true)
	println(resp.Text, err)
}

func Special(r *zdpgo_requests.Requests, url string) {
	resp, err := r.Get(url)
	println(resp.Text, err)
	resp, err = r.Post(url)
	println(resp.Text, err)
	resp, err = r.Put(url)
	println(resp.Text, err)
	resp, err = r.Delete(url)
	println(resp.Text, err)
	resp, err = r.Patch(url)
	println(resp.Text, err)
}

func Proxy(r *zdpgo_requests.Requests, url, proxyUrl string) {
	r.SetProxy(proxyUrl)
	resp, err := r.Get(url)
	println(resp.Text, err)
	resp, err = r.Post(url)
	println(resp.Text, err)
	resp, err = r.Put(url)
	println(resp.Text, err)
	resp, err = r.Delete(url)
	println(resp.Text, err)
	resp, err = r.Patch(url)
	println(resp.Text, err)
	r.RemoveProxy()
}
func Params(r *zdpgo_requests.Requests, url string) {
	param := zdpgo_requests.Param{
		"a": "11",
		"b": "22.222",
		"c": "abc",
	}
	resp, err := r.Get(url, param)
	println(resp.Text, err)
	resp, err = r.Post(url, param)
	println(resp.Text, err)
	resp, err = r.Put(url, param)
	println(resp.Text, err)
	resp, err = r.Delete(url, param)
	println(resp.Text, err)
	resp, err = r.Patch(url, param)
	println(resp.Text, err)
}

// UploadFS 上传fs内嵌系统文件
func UploadFS(r *zdpgo_requests.Requests, url string, fsObj embed.FS) {
	r.IsFs = true
	r.Fs = fsObj
	fileMap := zdpgo_requests.Files{
		"file": "test/test.txt",
	}
	resp, err := r.Post(url, fileMap)
	println(resp.Text, err)
}

// Upload 上传普通文件
func Upload(r *zdpgo_requests.Requests, url string) {
	fileMap := zdpgo_requests.Files{
		"file": "test/test1.txt",
	}
	resp, err := r.Post(url, fileMap)
	println(resp.Text, err)
}

// Header 设置请求头
func Header(r *zdpgo_requests.Requests, url string) {
	header := zdpgo_requests.Header{
		"User-Agent": "zdpgo_11111",
		"Abc-123":    "zdpgo_11111",
	}
	resp, err := r.Post(url, header)
	println(resp.Text, err)
}

// Json 发送json数据
func Json(r *zdpgo_requests.Requests, url string) {
	// 发送JsonMap
	jMap := zdpgo_requests.JsonMap{
		"User-Agent": "zdpgo_11111",
		"Abc-123":    "zdpgo_11111",
	}
	resp, err := r.Post(url, jMap)
	println(resp.Text, err)

	// 发送json字符串
	var jStr zdpgo_requests.JsonString = "{\"aaabbbcc\":1122233}"
	resp, err = r.Post(url, jStr)
	println(resp.Text, err)
}

// Text 发送纯文本数据
func Text(r *zdpgo_requests.Requests, url string) {
	// 发送json字符串
	var jStr = "{\"aaabbbcc\":1122233}"
	resp, err := r.Post(url, jStr)
	println(resp.Text, err)
}

// Redirect 重定向
func Redirect(r *zdpgo_requests.Requests, url string) {
	resp, _ := r.Get(url)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.IsRedirect)
	fmt.Println(resp.RedirectUrl)
	fmt.Println(resp.StartTime)
	fmt.Println(resp.EndTime)
	fmt.Println(resp.SpendTime)
	fmt.Println(resp.SpendTimeSeconds)
}

// Download 下载
func Download(r *zdpgo_requests.Requests) {
	// 下载为bytes
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	_, err := r.DownloadToBytes(imgUrl)
	if err != nil {
		r.Log.Error(err.Error())
	}
	r.Log.Debug("下载成功")

	// 下载到指定目录
	imgUrl = "https://www.twle.cn/static/i/img1.jpg"
	r.Download(imgUrl, "tmp")

	// 下载到临时目录
	data := []struct {
		Url       string
		NotResult string
	}{
		{"https://www.twle.cn/static/i/img1.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
	}

	for _, url := range data {
		result := r.DownloadToTmp(url.Url)
		if result == url.NotResult {
			r.Log.Error("下载错误：不是期望的值")
		}
	}

	// 下载并删除
	data1 := []struct {
		Url       string
		IsDeleted bool
	}{
		{"https://www.twle.cn/static/i/img1.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
	}

	for _, url := range data1 {
		result := r.DownloadToTmpAndReturnIsDeleted(url.Url, 10)
		if result == url.IsDeleted {
			r.Log.Error("下载错误：不是期望的值")
		}
	}

	// 批量下载
	data2 := [][]string{
		{"https://www.twle.cn/static/i/img1.jpg", "https://images3.alphacoders.com/120/1205462.jpg"},
	}

	for _, urls := range data2 {
		r.DownloadMany(urls, "tmp")
	}
}

func Timeout(r *zdpgo_requests.Requests, url string) {
	resp, err := r.Get(url)
	if resp != nil {
		println(resp.Text, err)
	}
}

func Auth(r *zdpgo_requests.Requests, url string, username, password string) {
	r.SetBasicAuth(username, password)
	resp, err := r.Get(url)
	println(resp.Text, err)
}
