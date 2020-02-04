package main

import (
	"context"
	"github.com/e421083458/gateway_demo/proxy/load_balance"
	"github.com/e421083458/gateway_demo/proxy/middleware"
	proxy2 "github.com/e421083458/gateway_demo/proxy/proxy"
	"github.com/e421083458/gateway_demo/proxy/tcp_proxy"
	"log"
	"net"
)

var (
	addr = "127.0.0.1:2002"
)

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler"))
}

func main() {
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:2003", "40")

	//tcpServ := tcp_proxy.TcpServer{Addr: addr, Handler: &tcpHandler{},}
	//tcpServ.ListenAndServe()

	proxy := proxy2.NewTcpLoadBalanceReverseProxy(&middleware.TcpSliceRouterContext{}, rb)
	log.Println("Starting tcpserver at " + addr)
	log.Fatal(tcp_proxy.ListenAndServe(addr, proxy))
}
