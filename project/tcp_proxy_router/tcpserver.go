package tcp_proxy_router

import (
	"context"
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/reverse_proxy"
	"github.com/e421083458/gateway_demo/project/tcp_proxy_middleware"
	"github.com/e421083458/gateway_demo/project/tcp_server"
	"log"
)

var (
	tcpServerList []*tcp_server.TcpServer
)

func TcpServerRun() {
	serviceList := dao.ServiceHandler.GetTcpServiceList()
	fmt.Println("tcp server list", serviceList)
	for _, service := range serviceList {
		go func(service *dao.ServiceDetail) {
			addr := fmt.Sprint(service.TcpRule.Port)
			rb, err := service.GetTcpLoadBalancer()
			if err != nil {
				fmt.Errorf("GetTcpLoadBalancer err:%v", err)
				return
			}
			router := tcp_proxy_middleware.NewTcpSliceRouter()
			router.Group("/").Use(
				tcp_proxy_middleware.TcpServerFlowCountMiddleware(),
				tcp_proxy_middleware.TcpServerFlowLimitMiddleware(),
				tcp_proxy_middleware.TcpClientFlowLimitMiddleware(),
				tcp_proxy_middleware.TcpWhiteIplistMiddleware(),
				tcp_proxy_middleware.TcpBlackIplistMiddleware())

			routerHandler := tcp_proxy_middleware.NewTcpSliceRouterHandler(func(c *tcp_proxy_middleware.TcpSliceRouterContext) tcp_server.TCPHandler {
				return reverse_proxy.NewTcpLoadBalanceReverseProxy(c, rb)
			}, router)

			baseCtx := context.Background()
			baseCtx = context.WithValue(baseCtx, "service_detail", service)
			tcpServ := &tcp_server.TcpServer{
				Addr:    ":" + addr,
				Handler: routerHandler,
				BaseCtx: baseCtx}
			tcpServerList = append(tcpServerList, tcpServ)
			log.Printf(" [INFO] tcp_proxy_run %s\n", ":"+addr)
			if err := tcpServ.ListenAndServe(); err != nil {
				log.Fatalf(" [ERROR] tcp_proxy_run %s err:%v\n", ":"+addr, err)
			}
		}(service)
	}
}

func TcpServerStop() {
	for _, tcpServer := range tcpServerList {
		if err := tcpServer.Close(); err != nil {
			log.Printf(" [ERROR] tcp_proxy_stop err:%v\n", err)
			continue
		}
		log.Printf(" [INFO] tcp_proxy_stop stopped\n")
	}
}
