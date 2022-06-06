package zdpgo_requests

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_password"
	"github.com/zhangdapeng520/zdpgo_random"
)

type Requests struct {
	ClientPort int                      // 源端口
	Config     *Config                  // 配置对象
	Log        *zdpgo_log.Log           // 日志对象
	File       *zdpgo_file.File         // 文件对象
	Json       *zdpgo_json.Json         // json处理对象
	Random     *zdpgo_random.Random     // 随机数据生成
	Password   *zdpgo_password.Password // 加密对象
}

func New() *Requests {
	return NewWithConfig(&Config{})
}

// NewWithConfig 通过配置创建Requests请求对象
func NewWithConfig(config *Config) *Requests {
	r := Requests{}

	// 创建日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_requests.log"
	}
	r.Log = zdpgo_log.NewWithDebug(config.Debug, config.LogFilePath)

	// 配置
	if config.ContentType == "" {
		config.ContentType = "multipart/form-data"
	}
	if config.UserAgent == "" {
		config.UserAgent = "ZDP-Go-Requests"
	}
	if config.Timeout == 0 {
		config.Timeout = 60 // 默认请求不超过1分钟
	}
	if config.TmpDir == "" {
		config.TmpDir = ".zdpgo_requests_tmp_files"
	}
	r.Config = config // 配置对象

	r.Json = zdpgo_json.New()                                                 // 实例化json对象
	r.File = zdpgo_file.NewWithConfig(zdpgo_file.Config{Debug: config.Debug}) // 实例化文件对象
	r.Random = zdpgo_random.New()                                             // 随机数据对象
	r.Password = zdpgo_password.NewWithConfig(&zdpgo_password.Config{
		Debug:       config.Debug,
		LogFilePath: config.LogFilePath,
		EccKey: zdpgo_password.Key{
			PrivateKey: config.Ecc.PrivateKey,
			PublicKey:  config.Ecc.PublicKey,
		},
	})
	return &r
}

// RemoveProxy 移除代理
func (r *Requests) RemoveProxy(client *http.Client) {
	// 设置代理
	client.Transport = &http.Transport{
		Proxy:           nil,                                                     // 设置代理
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.IsCheckHttps}, // 是否跳过证书校验
	}
	client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间
	r.Config.ProxyUrl = ""
}
