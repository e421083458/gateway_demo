package main

import (
	"github.com/e421083458/gateway_demo/proxy/load_balance"
	"github.com/e421083458/gateway_demo/proxy/middleware"
	proxy2 "github.com/e421083458/gateway_demo/proxy/proxy"
	"log"
	"net/http"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	//客户端注册：https://github.com/e421083458/gateway_demo/blob/master/demo/proxy/real_server_register/main.go
	mConf, err := load_balance.NewLoadBalanceZkConf("http://%s/base",
		"/rs_server",
		[]string{"127.0.0.1:2181"},
		map[string]string{"127.0.0.1:2003": "20"})
	if err != nil {
		panic(err)
	}
	rb := load_balance.LoadBanlanceFactorWithConf(load_balance.LbWeightRoundRobin, mConf)
	proxy := proxy2.NewLoadBalanceReverseProxy(&middleware.SliceRouterContext{}, rb)
	//proxy := NewMultipleHostsReverseProxy(rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
