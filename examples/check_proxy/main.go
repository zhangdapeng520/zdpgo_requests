package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_yaml"
)

type ProxyConfig struct {
	ProxyUrl  string `yaml:"proxy_url"`
	TargetUrl string `yaml:"target_url"`
}

func main() {
	r := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		Debug:   true,
		Timeout: 5,
	})

	// 读取配置中的代理
	y := zdpgo_yaml.New()
	var proxyConfig ProxyConfig
	err := y.ReadConfig("config.yaml", &proxyConfig)
	if err != nil {
		r.Log.Error("读取代理配置失败", "error", err)
		return
	}
	r.Config.ProxyUrl = proxyConfig.ProxyUrl

	// 使用代理发送请求
	response, err := r.Get(proxyConfig.TargetUrl)
	if err != nil {
		r.Log.Error("使用代理发送请求失败", "error", err)
		return
	}
	r.Log.Debug("使用代理发送请求成功", "response", response.Text)
}
