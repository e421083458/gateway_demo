package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, //探活时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}

	// 创建客户端
	client := &http.Client{
		Timeout:   time.Second * 30, //请求超时时间
		Transport: transport,
	}

	// 请求数据
	req, err := http.NewRequest("GET", "http://127.0.0.1:2002/bye", nil)
	if err != nil {
		panic(err)
	}
	//jwtToken, err := public.Encode("foo")
	//if err != nil {
	//	panic(err)
	//}
	token:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJ0ZXN0IiwiZXhwIjoxNTg1Mzc0MTQzLCJpc3MiOiJ0ZXN0In0.ySGLRlSMgOLBz_bRk7puLBHP0AuHjFqVKIgSxlnhTtI"
	req.Header.Set("Authorization", "Bearer " + token)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bds))
}
