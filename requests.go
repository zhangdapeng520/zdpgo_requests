package zdpgo_requests

import (
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_random"
)

type Requests struct {
	Request *Request             // 请求对象
	Config  *Config              // 配置对象
	Log     *zdpgo_log.Log       // 日志对象
	File    *zdpgo_file.File     // 文件对象
	Json    *zdpgo_json.Json     // json处理对象
	Random  *zdpgo_random.Random // 随机数据生成
}

func New() *Requests {
	return NewWithConfig(Config{})
}

// NewWithConfig 通过配置创建Requests请求对象
func NewWithConfig(config Config) *Requests {
	r := Requests{}

	// 创建日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_requests.log"
	}
	logConfig := zdpgo_log.Config{
		Debug:       config.Debug,
		OpenJsonLog: true,
		LogFilePath: config.LogFilePath,
	}
	if config.Debug {
		logConfig.IsShowConsole = true
	}
	r.Log = zdpgo_log.NewWithConfig(logConfig)

	// 配置
	r.Config = &config

	// 实例化json对象
	r.Json = zdpgo_json.New()

	// 实例化请求对象
	r.Request = NewRequest()

	// 实例化文件对象
	r.File = zdpgo_file.NewWithConfig(zdpgo_file.Config{Debug: config.Debug})

	// 随机数据对象
	r.Random = zdpgo_random.New()

	return &r
}
