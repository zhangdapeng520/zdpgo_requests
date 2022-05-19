package zdpgo_requests

import (
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_random"
	"net/http"
)

type Requests struct {
	Request *Request               // 请求对象
	HttpReq *http.Request          // http请求对象
	Header  *http.Header           // 请求头
	Client  *http.Client           // 请求客户端
	Cookies []*http.Cookie         // cookie
	Params  []map[string]string    // 请求参数
	Forms   []map[string]string    // form表单数据
	Body    string                 // 请求体内容
	Files   []map[string]string    // 文件列表
	JsonMap map[string]interface{} // JSON数据

	Config *Config              // 配置对象
	Log    *zdpgo_log.Log       // 日志对象
	File   *zdpgo_file.File     // 文件对象
	Json   *zdpgo_json.Json     // json处理对象
	Random *zdpgo_random.Random // 随机数据生成
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
	r.Config = &config // 配置对象

	r.HttpReq = r.GetHttpRequest()             // HTTP请求对象
	r.Client = r.GetHttpClient()               // HTTP客户端对象
	r.Header = new(http.Header)                // 请求头
	r.Params = make([]map[string]string, 0, 0) // 请求参数
	r.Forms = make([]map[string]string, 0, 0)  // 表单数据
	r.Files = make([]map[string]string, 0, 0)  // 文件列表
	r.JsonMap = make(map[string]interface{})   // JSON数据

	r.Json = zdpgo_json.New()                                                 // 实例化json对象
	r.Request = NewRequest()                                                  // 实例化请求对象
	r.File = zdpgo_file.NewWithConfig(zdpgo_file.Config{Debug: config.Debug}) // 实例化文件对象
	r.Random = zdpgo_random.New()                                             // 随机数据对象

	return &r
}
