package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		Debug: true,
	})

	targetUrl := "http://localhost:3333/ping"
	jsonUrl := "http://localhost:3333/json"
	formUrl := "http://localhost:3333/form"
	textUrl := "http://localhost:3333/text"

	// 发送GET请求
	response, err := r.Any("GET", targetUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Text)

	// 发送POST请求，携带json参数
	response, err = r.Any("POST", jsonUrl, "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// 发送PUT请求，携带json参数
	response, err = r.Any("PUT", jsonUrl, "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// 发送PATCH请求，携带json参数
	response, err = r.Any("PATCH", jsonUrl, "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// 发送DELETE请求，携带json参数
	response, err = r.Any("DELETE", jsonUrl, "{\"age\":22}")
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// POST提交表单数据
	form := zdpgo_requests.Form{
		"username": "zhangdapeng520",
		"password": "zhangdapeng520",
	}
	response, err = r.Any("POST", formUrl, form)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// POST提交JSON数据，普通的map数据就是JSON数据
	jsonMap := map[string]interface{}{
		"username": "zhangdapeng520",
		"password": "zhangdapeng520",
	}
	response, err = r.Any("POST", jsonUrl, jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// PUT提交JSON数据，结构体类型默认作为json数据处理
	jsonStuct := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "zhangdapeng520",
		Password: "zhangdapeng520",
	}
	response, err = r.Any("POST", jsonUrl, jsonStuct)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)

	// POST提交纯文本数据
	response, err = r.Any("POST", textUrl, r.GetText("abc 123 *&X 中国我爱你"))
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)
}
