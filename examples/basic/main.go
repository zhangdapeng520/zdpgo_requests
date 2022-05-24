package main

import (
	"basic/router"
	"embed"
	"github.com/zhangdapeng520/zdpgo_requests"
)

//go:embed test/*
var fsObj embed.FS

func main() {
	r := zdpgo_requests.NewWithConfig(zdpgo_requests.Config{
		Debug:   true,
		Timeout: 5,
	})

	var (
		url       = "http://localhost:3333/ping"
		jsonUrl   = "http://localhost:3333/json"
		textUrl   = "http://localhost:3333/text"
		uploadUrl = "http://localhost:3333/upload"
		authUrl   = "http://localhost:3333/admin"
		//proxyUrl    = "http://10.1.3.12:9999"
		redirectUrl = "https://www.baidu.com/link?url=IIZcBDQ9FSkK8wRluFkNAxjf4a7VDwHH0kFqGazjEAFGRDdnxe0HqQRdSocksxbbrpMjo7PTBeGjgnmf0aYOqN7ld6dXDBVO_jMYS16Yuy7CI5M_TMysMLpmFhF4CEjGjXOEYvjL_r9Hgz2-4jwsoa"
	)

	router.Any(r, url)                                    // any方法
	router.Special(r, url)                                // 特定方法
	router.Auth(r, authUrl, "zhangdapeng", "zhangdapeng") // 权限
	//router.Proxy(r, url, proxyUrl)                        // 代理
	router.Params(r, url)                // 查询参数
	router.UploadFS(r, uploadUrl, fsObj) // 上传嵌入系统的文件
	router.Upload(r, uploadUrl)          // 上传普通文件
	router.Header(r, url)                // 设置请求头
	router.Json(r, jsonUrl)              // 发送json数据
	router.Text(r, textUrl)              // 发送纯文本数据
	router.Redirect(r, redirectUrl)      // 重定向
	router.Download(r)                   // 下载

	router.Timeout(r, "http://localhost:3333/long") // 超时方法
}
