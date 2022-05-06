package zdpgo_requests

var VERSION string = "0.1.6"         // 版本编号
type Header map[string]string        // 请求头类型
type Params map[string]string        // Query查询参数类型
type Datas map[string]string         // POST提交的数据
type JsonData map[string]interface{} // 提交JSON格式的数据
type JsonString string               // 提交JSON格式的字符串
type Files map[string]string         // 文件列表：name ,filename
type Auth []string                   // Auth 权限校验类型，{username,password}