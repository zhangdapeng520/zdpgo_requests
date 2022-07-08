package zdpgo_requests

/*
@Time : 2022/4/28 17:47
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: requests请求相关配置
*/

type Config struct {
	PoolSize               int    `json:"pool_size" yaml:"pool_size"`                     // 最多同时执行任务数量，默认333
	LimitSleepSeconds      int    `json:"limit_sleep_seconds" yaml:"limit_sleep_seconds"` // 达到限制后的休眠时间，默认3秒
	Timeout                int    `json:"timeout" yaml:"timeout"`                         // 请求超时时间（秒）
	ContentType            string `yaml:"content_type" json:"content_type"`               // 内容类型，默认"application/json"
	UserAgent              string `yaml:"user_agent" json:"user_agent"`                   // 用户代理，默认"ZDP-Go-Requests"
	Author                 string `yaml:"author" json:"author"`                           // 作者，自定义请求头
	TmpDir                 string `yaml:"tmp_dir" json:"tmp_dir"`                         // 文件上传临时目录
	IsCheckHttps           bool   `json:"is_check_https" yaml:"is_check_https"`           // 是否检查HTTPS
	IsCheckRedirect        bool   `yaml:"is_check_redirect" json:"is_check_redirect"`
	TargetUrl              string `yaml:"target_url" json:"target_url"` // 目标地址
	IsRecordRequestDetail  bool   `yaml:"is_record_request_detail" json:"is_record_request_detail"`
	IsRecordResponseDetail bool   `yaml:"is_record_response_detail" json:"is_record_response_detail"`
	IsKeepSession          bool   `yaml:"is_keep_session" json:"is_keep_session"`
	IsRandomUserAgent      bool   `yaml:"is_random_user_agent" json:"is_random_user_agent"` // 随机的用户代理
	ProxyUrl               string `yaml:"is_json" json:"is_json"`
}
