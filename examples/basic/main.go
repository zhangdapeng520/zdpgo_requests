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
		uploadUrl = "http://localhost:3333/upload"
		authUrl   = "http://localhost:3333/admin"
		proxyUrl  = "http://10.1.3.12:9999"
	)

	router.Any(r, url)                                    // any方法
	router.Special(r, url)                                // 特定方法
	router.Auth(r, authUrl, "zhangdapeng", "zhangdapeng") // 权限
	router.Proxy(r, url, proxyUrl)                        // 代理
	router.Params(r, url)                                 // 查询参数
	router.UploadFS(r, uploadUrl, fsObj)

	router.Timeout(r, "http://localhost:3333/long") // 超时方法
}
