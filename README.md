# zdpgo_requests
Golang中用于发送HTTP请求的库

## 版本历史
- v0.1.0 2022/4/9   新增GET和POST请求
- v0.1.1 2022/4/11  POST的map默认当表单数据
- v0.1.2 2022/4/11  添加忽略URL解析错误的请求方法
- v0.1.3 2022/4/12  支持POST纯文本数据
- v0.1.4 2022/4/12  代码重构
- v0.1.5 2022/4/13  支持任意类型HTTP请求
- v0.1.6 2022/4/13  支持设置代理
- v0.1.7 2022/4/13  支持发送JSON数据
- v0.1.8 2022/4/16  解决部分URL无法正常请求的BUG
- v0.1.9 2022/4/18  BUG修复：header请求头重复
- v0.2.0 2022/4/18  新增：获取请求和响应详情
- v0.2.1 2022/4/20  新增：获取响应状态码
- v0.2.2 2022/4/20  新增：下载文件
- v0.2.3 2022/4/21  新增：文件上传
- v0.2.4 2022/4/22  新增：支持上传FS文件系统文件
- v0.2.5 2022/4/28  新增：检查重定向和请求消耗时间
- v0.2.6 2022/5/6   新增：根据字节数组上传文件
- v0.2.7 2022/5/8   新增：根据超时时间发送POST请求并携带JSON数据
  
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
	r := zdpgo_New()
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
	var jsonStr Datas = map[string]string{
		"username": "zhangdapeng520",
	}
	var headers Header = map[string]string{
		"Content-Type": "application/json",
	}
	resp, _ = Post(url, true, jsonStr, headers)
	println(resp.Text())

	// 权限校验
	resp, _ = r.Get(
		url,
		Auth{"zhangdapeng520", "password...."},
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
	r := zdpgo_New()
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
	req := zdpgo_New()

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
	r := zdpgo_New()

	// 发送JSON字符串
	var jsonStr JsonString = "{\"name\":\"requests_post_test\"}"
	resp, _ := r.Post("http://localhost:8888", jsonStr)
	println(resp.Text())

	// 发送map
	var data JsonData = make(map[string]interface{})
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

### 设置请求头
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	// 直接设置请求头
	req := zdpgo_New()
	req.Request.Header.Set("accept-encoding", "gzip, deflate, br")
	resp, _ := req.Get("http://localhost:8888", false, Header{"Referer": "http://127.0.0.1:9999"})
	println(resp.Text())

	// 将请求头作为参数传递
	h := Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}
	h2 := Header{
		"Referer":         "http://localhost:8888",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"User-Agent":      "zdpgo_requests",
	}
	resp, _ = req.Get("http://localhost:8888", h, h2)
	println(resp.Text())
}
```

### 设置查询参数
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	req := zdpgo_New()
	p := Params{
		"name": "file",
		"id":   "12345",
	}
	resp, _ := req.Get("http://localhost:8888", false, p)
	fmt.Println(resp.Text())
}
```

### Basic Auth权限校验
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	req := zdpgo_New()
	resp, _ := req.Get(
		"http://localhost:8080/admin/secrets",
		Auth{"zhangdapeng", "zhangdapeng"},
	)
	println(resp.Text())
}
```

### 获取请求和响应详情
```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
	"github.com/zhangdapeng520/zdpgo_requests/core/requests"
)

func main() {
	// 发送GET请求
	r := zdpgo_New()
	baseUrl := "http://10.1.3.12:8888/"
	query := "?a=<script>alert(\"XSS\");</script>&b=UNION SELECT ALL FROM information_schema AND ' or SLEEP(5) or '&c=../../../../etc/passwd"
	url := baseUrl + query

	var h1 Header = Header{"a": "111", "b": "222"}
	resp, err := r.GetIgnoreParseError(url, h1)
	if err != nil {
		fmt.Println("错误2", err)
	} else {
		println(resp.Text())
		println("请求详情：\n", resp.RawReqDetail)
		println("响应详情：\n", resp.RawRespDetail)
	}

	var h2 Header = Header{"c": "333", "d": "444"}
	resp1, err := r.GetIgnoreParseError(url, h2)
	if err != nil {
		fmt.Println("错误3", err)
	} else {
		println(resp1.Text())
		println("请求详情：\n", resp.RawReqDetail)
		println("响应详情：\n", resp.RawRespDetail)
	}
}
```

### 下载图片
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_New()
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	err := r.Download(imgUrl, "test1.jpg")
	if err != nil {
		panic(err)
	}
}
```

### 文件上传
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func main() {
	r := zdpgo_New()
	imgUrl := "http://localhost:8888/upload"
	err := r.Upload(imgUrl, "test1.jpg")
	if err != nil {
		panic(err)
	}
}
```

### 上传FS文件系统文件
```go
package main

import (
	"embed"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

//go:embed test/*
var fsObj embed.FS

func main() {
	r := zdpgo_New()

	targetUrl := "http://localhost:8888/upload"
	filename := "test/main.go"

	respBytes, err := r.UploadFsToBytes(targetUrl, fsObj, "file", filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(respBytes))

	respString, err := r.UploadFsToString(targetUrl, fsObj, "file", filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(respString)
}
```