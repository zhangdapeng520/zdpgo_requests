package zdpgo_requests

import (
	"bytes"
)

type Header map[string]string       // 请求头类型
type Param map[string]string        // Query查询参数类型
type Form map[string]string         // POST提交的数据
type JsonMap map[string]interface{} // 提交JSON格式的数据
type JsonString string              // 提交JSON格式的字符串
type Files map[string]string        // 文件列表：name ,filename

// FormFileBytes 字节类型的表单文件
type FormFileBytes struct {
	FormName     string `json:"form_name"`     // 表单名称
	FileName     string `json:"file_name"`     // 文件名称
	ContentBytes []byte `json:"content_bytes"` // 文件内容
}

// Request 请求对象
type Request struct {
	Method string            `json:"method"`
	Url    string            `json:"url"`
	Header map[string]string `json:"header"`
	Body   *bytes.Buffer     `json:"body"`
}

// BaseAuth 基础权限校验类型
type BaseAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response 响应对象
type Response struct {
	Content          []byte `json:"content"`            // 响应内容
	Text             string `json:"text"`               // 响应文本
	RawReqDetail     string `json:"raw_req_detail"`     // 请求详情字符串
	RawRespDetail    string `json:"raw_resp_detail"`    // 响应详情字符串
	StatusCode       int    `json:"status_code"`        // 状态码
	IsRedirect       bool   `json:"is_redirect"`        // 是否重定向了
	RedirectUrl      string `json:"redirect_url"`       // 重定向的的URL地址
	StartTime        int    `json:"start_time"`         // 请求开始时间（纳秒）
	EndTime          int    `json:"end_time"`           // 请求结束时间（纳秒）
	SpendTime        int    `json:"spend_time"`         // 请求消耗时间（纳秒）
	SpendTimeSeconds int    `json:"spend_time_seconds"` // 请求消耗时间（秒）
	ClientIp         string `json:"client_ip"`          // 客户端IP
	ClientPort       int    `json:"client_port"`        // 客户端端口号
}
