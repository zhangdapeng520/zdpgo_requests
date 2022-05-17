package zdpgo_requests

import "github.com/zhangdapeng520/zdpgo_log"

/*
@Time : 2022/5/17 17:22
@Author : 张大鹏
@File : logger.go
@Software: Goland2021.3.1
@Description: logger 日志相关
*/

var Log *zdpgo_log.Log // 日志对象

func init() {
	if Log == nil {
		Log = zdpgo_log.NewWithConfig(zdpgo_log.Config{
			Debug:         true,
			IsShowConsole: true,
			OpenJsonLog:   true,
			LogFilePath:   "logs/zdpgo/zdpgo_requests.log",
		})
	}
}
