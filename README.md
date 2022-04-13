# zdpgo_requests
Golang中用于发送HTTP请求的库

## 版本历史
- 版本0.1.0 2022年4月9日 新增GET和POST请求
- 版本0.1.1 2022年4月11日 POST的map默认当表单数据
- 版本0.1.2 2022年4月11日 添加忽略URL解析错误的请求方法
- 版本0.1.3 2022年4月12日 支持POST纯文本数据
- 版本0.1.4 2022年4月12日 代码重构
- 版本0.1.5 2022年4月13日 支持任意类型HTTP请求
- 版本0.1.6 2022年4月13日 支持设置代理
- 版本0.1.7 2022年4月13日 支持发送JSON数据

## 使用案例
### 快速入门
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_requests.New()
	url := "http://localhost:8888"
	resp, err := r.Get(url, true)
	if err != nil {
		return
	}
	println(resp.Text())

	// 发送POST请求
	data := map[string]string{
		"name": "request123",
	}
	resp, _ = r.Post(url, data)
	println(resp.Text())

	// 发送json数据
	var jsonStr requests.Datas = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers requests.Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ = requests.Post(url, true, jsonStr, headers)
	println(resp.Text())

	// 权限校验
	resp, _ = r.Get(
		url,
		requests.Auth{"zhangdapeng520", "password...."},
	)
	println(resp.Text())
}
```

### 发送不同类型的请求
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_requests.New()
	targetUrl := "http://localhost:8888"

	// 发送GET请求
	resp, _ := r.Get(targetUrl)
	println(resp.Text())

	// 发送POST请求
	resp, _ = r.Post(targetUrl)
	println(resp.Text())

	// 发送PUT请求
	resp, _ = r.Put(targetUrl)
	println(resp.Text())

	// 发送DELETE请求
	resp, _ = r.Delete(targetUrl)
	println(resp.Text())

	// 发送PATCH请求
	resp, _ = r.Patch(targetUrl)
	println(resp.Text())
}
```

### 设置代理
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	req := zdpgo_requests.New()

	// 设置代理
	err := req.SetProxy("http://localhost:8888")
	if err != nil {
		panic(err)
	}

	// 发送请求
	// 设置了代理以后，请求被重定向了代理的URL
	resp, _ := req.Get("http://localhost:9999", false)
	fmt.Println("响应：", resp.Text())
}
```

### 发送JSON数据
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	r := zdpgo_requests.New()

	// 发送JSON字符串
	var jsonStr requests.JsonString = "{\"name\":\"requests_post_test\"}"
	resp, _ := r.Post("http://localhost:8888", jsonStr)
	println(resp.Text())

	// 发送map
	var data requests.JsonData = make(map[string]interface{})
	data["name"] = "root"
	data["password"] = "root"
	data["host"] = "localhost"
	data["port"] = 8888
	resp, _ = r.Post("http://localhost:8888", data)
	println(resp.Text())

	// PUT发送JSON数据
	resp, _ = r.Put("http://localhost:8888", data)
	println(resp.Text())

	// DELETE发送JSON数据
	resp, _ = r.Delete("http://localhost:8888", data)
	println(resp.Text())

	// PATCH发送JSON数据
	resp, _ = r.Patch("http://localhost:8888", data)
	println(resp.Text())

	// 发送纯文本数据（非json）
	resp, _ = r.Post("http://localhost:8888", "纯文本内容。。。")
	println(resp.Text())
}
```