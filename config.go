package zdpgo_requests

/*
@Time : 2022/4/28 17:47
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: requests请求相关配置
*/

type Config struct {
	Timeout       int  `env:"timeout" json:"timeout" yaml:"timeout"`                      // 请求超时时间（秒）
	CheckRedirect bool `env:"check_redirect" json:"check_redirect" yaml:"check_redirect"` // 是否检查重定向
}
