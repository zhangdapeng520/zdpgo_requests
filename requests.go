package zdpgo_requests

import (
	"crypto/tls"
	"embed"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_random"
	"net/http"
	"time"
)

type Requests struct {
	HttpReq      *http.Request          // http请求对象
	HttpResponse *http.Response         // http响应对象
	Response     *Response              // 自定义响应对象
	Header       *http.Header           // 请求头
	Client       *http.Client           // 请求客户端
	Cookies      []*http.Cookie         // cookie
	Params       []map[string]string    // 请求参数
	Forms        []map[string]string    // form表单数据
	Body         string                 // 请求体内容
	Files        []map[string]string    // 文件列表
	JsonMap      map[string]interface{} // JSON数据
	Fs           embed.FS               // 嵌入文件系统
	IsFs         bool                   // 是否使用嵌入文件系统

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
	if config.TmpDir == "" {
		config.TmpDir = ".zdpgo_requests_tmp_files"
	}
	if config.ClientPort == 0 {
		config.ClientPort = 33334
	}
	r.Config = &config // 配置对象

	r.Json = zdpgo_json.New()                                                 // 实例化json对象
	r.File = zdpgo_file.NewWithConfig(zdpgo_file.Config{Debug: config.Debug}) // 实例化文件对象
	r.Random = zdpgo_random.New()                                             // 随机数据对象

	return &r
}

// RemoveProxy 移除代理
func (r *Requests) RemoveProxy() {
	// 设置代理
	r.Client.Transport = &http.Transport{
		Proxy:           nil,                                                   // 设置代理
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.CheckHttps}, // 是否跳过证书校验
	}
	r.Client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间
}

// InitData 初始化数据
func (r *Requests) InitData() {
	r.HttpResponse = &http.Response{} // HTTP响应对象
	r.Response = &Response{}          // 自定义响应对象
	r.Header = &http.Header{}         // 请求头
	r.Header.Set("Content-Type", r.Config.ContentType)
	r.Header.Set("User-Agent", r.Config.UserAgent)

	r.Params = make([]map[string]string, 0, 0) // 请求参数
	r.Forms = make([]map[string]string, 0, 0)  // 表单数据
	r.Files = make([]map[string]string, 0, 0)  // 文件列表
	r.JsonMap = make(map[string]interface{})   // JSON数据

	r.HttpReq = r.GetHttpRequest() // HTTP请求对象
	r.Client = r.GetHttpClient()   // HTTP客户端对象

	// 处理fs文件系统
	r.IsFs = false
}

// SetBasicAuth 设置基本认证信息
func (r *Requests) SetBasicAuth(username, password string) {
	if r.HttpReq != nil {
		r.HttpReq.SetBasicAuth(username, password)
	}
}
