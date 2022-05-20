package zdpgo_requests

/*
@Time : 2022/4/28 17:47
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: requests请求相关配置
*/

type Config struct {
	Debug                  bool   `yaml:"debug" json:"debug"`                   // 是否为DEBUG模式
	LogFilePath            string `yaml:"log_file_path" json:"log_file_path"`   // 日志存放路径
	Timeout                int    `json:"timeout" yaml:"timeout"`               // 请求超时时间（秒）
	CheckRedirect          bool   `json:"check_redirect" yaml:"check_redirect"` // 是否检查重定向
	CheckHttps             bool   `json:"check_https" yaml:"check_https"`       // 是否检查HTTPS
	ContentType            string `yaml:"content_type" json:"content_type"`     // 内容类型，默认"multipart/form-data"
	UserAgent              string `yaml:"user_agent" json:"user_agent"`         // 用户代理，默认"ZDP-Go-Requests"
	TmpDir                 string `yaml:"tmp_dir" json:"tmp_dir"`               // 文件上传临时目录
	IsIgnoredParsedError   bool   `yaml:"is_ignored_parsed_error" json:"is_ignored_parsed_error"`
	TargetUrl              string `yaml:"target_url" json:"target_url"` // 目标地址
	IsRecordRequestDetail  bool   `yaml:"is_record_request_detail" json:"is_record_request_detail"`
	IsRecordResponseDetail bool   `yaml:"is_record_response_detail" json:"is_record_response_detail"`
	IsKeepSession          bool   `yaml:"is_keep_session" json:"is_keep_session"`
}
