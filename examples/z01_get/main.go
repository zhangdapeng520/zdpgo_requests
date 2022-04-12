package main

import (
	"github.com/zhangdapeng520/zdpgo_requests"
)

func demo1() {
	r := zdpgo_requests.New()
	url := "http://www.zhanluejia.net.cn"
	resp, err := r.Get(url, true)
	if err != nil {
		return
	}
	println(resp.Text())
}

func demo2() {
	r := zdpgo_requests.New()
	url := "http://localhost:9999"
	resp, err := r.Get(url, true)
	if err != nil {
		return
	}
	println(resp.Text())
}

func demo3() {
	r := zdpgo_requests.New()
	//resp, err := r.Get("http://10.1.3.12:8888/file=?%25%2532%2565%25%2532%2565%25%2532%2566%25%2532%2565%25%2532%2565%25%2532%2566%25%2532%2565%25%2532%2565%25%2532%2566%25%2532%2565%25%2532%2565%25%2532%2566%25%2532%2565%25%2532%2565%25%2532%2566%25%2536%2535%25%2537%2534%25%2536%2533%25%2532%2566%25%2537%2530%25%2536%2531%25%2537%2533%25%2537%2533%25%2537%2537%25%2536%2534")
	url := "http://10.1.3.12:8888/file=%%35%63%%32%65%%32%65%%35%63%%32%65%%32%65%%35%63%%3\n2%65%%32%65%%35%63%%32%65%%32%65%%35%63%%32%65%%32%65%%35%63%%32%65%%32%65%%35%6\n3%%35%37%%34%39%%34%65%%34%34%%34%66%%35%37%%35%33%%35%63%%37%37%%36%39%%36%65%%\n32%65%%36%39%%36%65%%36%39"
	resp, err := r.GetIgnoreParseError(url, true)
	if err != nil {
		return
	}
	println(resp.Text())
}

func main() {
	//demo2()
	demo3()
}
