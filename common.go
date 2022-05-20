package zdpgo_requests

type Header map[string]string       // 请求头类型
type Param map[string]string        // Query查询参数类型
type Form map[string]string         // POST提交的数据
type JsonMap map[string]interface{} // 提交JSON格式的数据
type JsonString string              // 提交JSON格式的字符串
type Files map[string]string        // 文件列表：name ,filename

// BaseAuth 基础权限校验类型
type BaseAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response 响应对象
type Response struct {
	Content          []byte // 响应内容
	Text             string // 响应文本
	RawReqDetail     string // 请求详情字符串
	RawRespDetail    string // 响应详情字符串
	StatusCode       int    // 状态码
	IsRedirect       bool   // 是否重定向了
	RedirectUrl      string // 重定向的的URL地址
	StartTime        int    // 请求开始时间（纳秒）
	EndTime          int    // 请求结束时间（纳秒）
	SpendTime        int    // 请求消耗时间（纳秒）
	SpendTimeSeconds int    // 请求消耗时间（秒）
}
