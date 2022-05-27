package zdpgo_requests

import (
	"fmt"
	"testing"
)

var (
	urlPath     = "http://localhost:3333/ping"
	proxyUrl    = "http://10.1.3.12:8080"
	uploadUrl   = "http://localhost:3333/upload"
	jsonUrl     = "http://localhost:3333/json"
	textUrl     = "http://localhost:3333/text"
	redirectUrl = "https://www.baidu.com/link?url=IIZcBDQ9FSkK8wRluFkNAxjf4a7VDwHH0kFqGazjEAFGRDdnxe0HqQRdSocksxbbrpMjo7PTBeGjgnmf0aYOqN7ld6dXDBVO_jMYS16Yuy7CI5M_TMysMLpmFhF4CEjGjXOEYvjL_r9Hgz2-4jwsoa"
	timeoutUrl  = "http://localhost:3333/long"
	authUrl     = "http://localhost:3333/admin"
	r           = getRequests()
)

func getRequests() *Requests {
	r := NewWithConfig(Config{
		Debug:    true,
		Timeout:  5,
		ProxyUrl: proxyUrl,
	})
	return r
}

func TestRequests_Any(t *testing.T) {
	resp, err := r.Any("get", urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Any("post", urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Any("put", urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Any("delete", urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Any("patch", urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// 任意类型的方法，不解析URL路径
func TestRequests_AnyNoParse(t *testing.T) {
	data := []Request{
		{"GET", urlPath, nil, nil, nil, nil},
		{"GET", urlPath, nil, nil, map[string]string{"a": "b"}, nil},
		{"POST", urlPath, nil, nil, nil, nil},
		{"DELETE", urlPath, nil, nil, nil, nil},
		{"PUT", urlPath, nil, nil, nil, nil},
		{"PATCH", urlPath, nil, nil, nil, nil},
	}

	for _, request := range data {
		response, err := r.AnyNoParseURL(request)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("状态码不是200")
		}
	}
}

func TestRequests_AnyJson(t *testing.T) {
	jsonData := map[string]interface{}{"a": 1, "b": 2.2, "c": "33", "d": true}
	data := []Request{
		{"POST", jsonUrl, nil, nil, nil, jsonData},
		{"DELETE", jsonUrl, nil, nil, nil, jsonData},
		{"PUT", jsonUrl, nil, nil, nil, jsonData},
		{"PATCH", jsonUrl, nil, nil, nil, jsonData},
	}

	for _, request := range data {
		response, err := r.AnyJson(request)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("状态码不是200")
		}
	}
}

func TestRequests_Special(t *testing.T) {
	resp, err := r.Get(urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Post(urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Put(urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Delete(urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Patch(urlPath)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

func TestRequests_Params(t *testing.T) {
	param := Param{
		"a": "11",
		"b": "22.222",
		"c": "abc",
	}
	resp, err := r.Get(urlPath, param)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Post(urlPath, param)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Put(urlPath, param)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Delete(urlPath, param)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	resp, err = r.Patch(urlPath, param)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// UploadFS 上传fs内嵌系统文件
func TestRequests_UploadFS(t *testing.T) {
	r.IsFs = true
	r.Fs = fsObj
	fileMap := Files{
		"file": "test/test.txt",
	}
	resp, err := r.Post(uploadUrl, fileMap)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// Upload 上传普通文件
func TestRequests_Upload(t *testing.T) {
	fileMap := Files{
		"file": "test/test1.txt",
	}
	resp, err := r.Post(uploadUrl, fileMap)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// Header 设置请求头
func TestRequests_Header(t *testing.T) {
	header := Header{
		"User-Agent": "zdpgo_11111",
		"Abc-123":    "zdpgo_11111",
	}
	resp, err := r.Post(urlPath, header)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// Json 发送json数据
func TestRequests_Json(t *testing.T) {
	// 发送JsonMap
	jMap := JsonMap{
		"User-Agent": "zdpgo_11111",
		"Abc-123":    "zdpgo_11111",
	}
	resp, err := r.Post(jsonUrl, jMap)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)

	// 发送json字符串
	var jStr JsonString = "{\"aaabbbcc\":1122233}"
	resp, err = r.Post(jsonUrl, jStr)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// Text 发送纯文本数据
func TestRequests_Text(t *testing.T) {
	// 发送json字符串
	var jStr = "{\"aaabbbcc\":1122233}"
	resp, err := r.Post(textUrl, jStr)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}

// Redirect 重定向
func TestRequests_Redirect(t *testing.T) {
	resp, _ := r.Get(redirectUrl)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.IsRedirect)
	fmt.Println(resp.RedirectUrl)
	fmt.Println(resp.StartTime)
	fmt.Println(resp.EndTime)
	fmt.Println(resp.SpendTime)
	fmt.Println(resp.SpendTimeSeconds)
}

// Download 下载
func TestRequests_Download(t *testing.T) {
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

func TestRequests_Timeout(t *testing.T) {
	r.SetTimeout(1) // 设置超时
	resp, err := r.Get(timeoutUrl)
	if resp != nil {
		r.Log.Debug("发送请求成功", "resp", resp, "error", err)
	}
}

func TestRequests_Auth(t *testing.T) {
	r.SetBasicAuth("zhangdapeng", "zhangdapeng")
	resp, err := r.Get(authUrl)
	r.Log.Debug("发送请求成功", "resp", resp, "error", err)
}
