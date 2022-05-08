package zdpgo_requests

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/8 19:05
@Author : 张大鹏
@File : post_test.go
@Software: Goland2021.3.1
@Description: 测试post相关的方法
*/

func TestRequests_PostJsonWithTimeout(t *testing.T) {
	r := getRequests()
	data := make(map[string]interface{})
	data["a"] = 111
	data["b"] = 222.222
	result, err := r.PostJsonWithTimeout("http://localhost:8889/payload", data, 1000)
	fmt.Println(result, err)
}
