package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New(zdpgo_log.NewWithDebug(true, "log.log"))
	targetUrl := "http://localhost:3333/ping"
	jsonUrl := "http://localhost:3333/json"
	textUrl := "http://localhost:3333/text"

	// 发送GET请求
	response, err := r.AnyCompareStatusCode("GET", targetUrl, targetUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// 发送POST请求，携带json参数
	response, err = r.AnyCompareStatusCode("POST", jsonUrl, jsonUrl+"/abcdefg", "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// 发送PUT请求，携带json参数
	response, err = r.AnyCompareStatusCode("PUT", jsonUrl, jsonUrl+"/abcdefg", "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// 发送PATCH请求，携带json参数
	response, err = r.AnyCompareStatusCode("PATCH", jsonUrl, jsonUrl+"/abcdefg", "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// 发送DELETE请求，携带json参数
	response, err = r.AnyCompareStatusCode("DELETE", jsonUrl, jsonUrl+"/abcdefg", "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// POST提交表单数据
	form := zdpgo_requests.Form{
		"username": "zhangdapeng520",
		"password": "zhangdapeng520",
	}
	response, err = r.AnyCompareStatusCode("POST", jsonUrl, jsonUrl+"/abcdefg", form)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// POST提交JSON数据，普通的map数据就是JSON数据
	jsonMap := map[string]interface{}{
		"username": "zhangdapeng520",
		"password": "zhangdapeng520",
	}
	response, err = r.AnyCompareStatusCode("POST", jsonUrl, jsonUrl+"/abcdefg", jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// PUT提交JSON数据，结构体类型默认作为json数据处理
	jsonStuct := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "zhangdapeng520",
		Password: "zhangdapeng520",
	}
	response, err = r.AnyCompareStatusCode("POST", jsonUrl, jsonUrl+"/abcdefg", jsonStuct)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)

	// POST提交纯文本数据
	response, err = r.AnyCompareStatusCode("POST", textUrl, textUrl, r.GetText("abc 123 *&X 中国我爱你"))
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.FirstStatusCode)
}
