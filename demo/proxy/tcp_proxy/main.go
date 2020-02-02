package main

import (
	"github.com/e421083458/gateway_demo/proxy/load_balance"
	"github.com/e421083458/gateway_demo/proxy/middleware"
	"github.com/e421083458/gateway_demo/proxy/tcp_proxy"
	"log"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:2003", "40")
	proxy := tcp_proxy.NewLoadBalanceReverseProxy(&middleware.TcpSliceRouterContext{}, rb)
	log.Println("Starting tcpserver at " + addr)
	log.Fatal(tcp_proxy.ListenAndServe(addr, proxy))
}
