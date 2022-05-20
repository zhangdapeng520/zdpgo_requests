package zdpgo_requests

import (
	"crypto/tls"
	"embed"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_random"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Requests struct {
	HttpReq *http.Request          // http请求对象
	Header  *http.Header           // 请求头
	Client  *http.Client           // 请求客户端
	Cookies []*http.Cookie         // cookie
	Params  []map[string]string    // 请求参数
	Forms   []map[string]string    // form表单数据
	Body    string                 // 请求体内容
	Files   []map[string]string    // 文件列表
	JsonMap map[string]interface{} // JSON数据
	Fs      embed.FS               // 嵌入文件系统
	IsFs    bool                   // 是否使用嵌入文件系统

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
	if config.FsTmpDir == "" {
		config.FsTmpDir = "zdpgo_requests_tmp_uploads"
	}
	r.Config = &config // 配置对象

	r.InitData()                   // 初始化数据
	r.HttpReq = r.GetHttpRequest() // HTTP请求对象
	r.Client = r.GetHttpClient()   // HTTP客户端对象

	r.Json = zdpgo_json.New()                                                 // 实例化json对象
	r.File = zdpgo_file.NewWithConfig(zdpgo_file.Config{Debug: config.Debug}) // 实例化文件对象
	r.Random = zdpgo_random.New()                                             // 随机数据对象

	return &r
}

// SetProxy 设置代理
func (r *Requests) SetProxy(proxyUrl string) error {
	// 解析代理地址
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		r.Log.Error("解析代理地址失败", "error", err, "proxyUrl", proxyUrl)
	}

	// 设置代理
	r.Client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(uri),                                    // 设置代理
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !r.Config.CheckHttps}, // 是否跳过证书校验
	}
	r.Client.Timeout = time.Second * time.Duration(r.Config.Timeout) // 超时时间

	return nil
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

// Exists 判断文件是否存在
func (r *Requests) Exists(filePath string) bool {
	// 判断文件是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

// DeleteDir 删除文件夹
func (r *Requests) DeleteDir(dirPath string) {
	if r.Exists(dirPath) {
		// 删除
		var (
			count = 0
			err   error
		)
		for count < 3 {
			err = os.RemoveAll(dirPath)
			count++
			time.Sleep(time.Second)
		}
		if err != nil {
			r.Log.Error("删除文件夹失败", "error", err, "dir", dirPath)
		}
	}
}

// InitData 初始化数据
func (r *Requests) InitData() {
	r.Header = &http.Header{} // 请求头
	r.Header.Set("Content-Type", r.Config.ContentType)
	r.Header.Set("User-Agent", r.Config.UserAgent)

	r.Params = make([]map[string]string, 0, 0) // 请求参数
	r.Forms = make([]map[string]string, 0, 0)  // 表单数据
	r.Files = make([]map[string]string, 0, 0)  // 文件列表
	r.JsonMap = make(map[string]interface{})   // JSON数据

	// 处理HTTP
	if r.HttpReq != nil {
		r.HttpReq.Body = nil        // 清空请求体
		r.HttpReq.GetBody = nil     // 清空get参数
		r.HttpReq.ContentLength = 0 // 清空内容长度
	}

	// 处理fs文件系统
	r.IsFs = false
}

// SetBasicAuth 设置基本认证信息
func (r *Requests) SetBasicAuth(username, password string) {
	if r.HttpReq != nil {
		r.HttpReq.SetBasicAuth(username, password)
	}
}
