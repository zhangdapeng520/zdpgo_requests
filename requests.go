package zdpgo_requests

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"

	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_password"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Requests struct {
	ClientPort int                      // 源端口
	Config     *Config                  // 配置对象
	File       *zdpgo_file.File         // 文件对象
	Json       *zdpgo_json.Json         // json处理对象
	Password   *zdpgo_password.Password // 加密对象
	TaskNum    int                      // 任务数量
}

func New() *Requests {
	return NewWithConfig(&Config{})
}

// NewWithConfig 通过配置创建Requests请求对象
func NewWithConfig(config *Config) *Requests {
	r := &Requests{}

	// 配置
	if config.ContentType == "" {
		config.ContentType = "application/json"
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
	if config.PoolSize == 0 {
		config.PoolSize = 333
	}
	if config.LimitSleepSeconds == 0 {
		config.LimitSleepSeconds = 3
	}
	r.Config = config         // 配置对象
	r.Json = zdpgo_json.New() // 实例化json对象
	r.File = zdpgo_file.New() // 实例化文件对象
	r.Password = zdpgo_password.New()

	// 返回
	return r
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

// GetRandomUserAgent 获取随机的用户代理
func (r *Requests) GetRandomUserAgent() string {
	var (
		length = len(CAgents)
		index  = rand.Intn(length)
	)
	return CAgents[index]
}
