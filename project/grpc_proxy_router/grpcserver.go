package grpc_proxy_router

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/grpc_proxy_middleware"
	"github.com/e421083458/gateway_demo/project/reverse_proxy"
	"github.com/e421083458/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcServerList []*grpc.Server
)

func GrpcServerRun() {
	serviceList := dao.ServiceHandler.GetGrpcServiceList()
	fmt.Println("grpc server list", serviceList)
	for _, service := range serviceList {
		go func(service *dao.ServiceDetail) {
			addr := fmt.Sprint(service.GrpcRule.Port)
			rb, err := service.GetTcpLoadBalancer()
			if err != nil {
				fmt.Errorf("GetGrpcLoadBalancer err:%v", err)
				return
			}

			lis, err := net.Listen("tcp", ":"+addr)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			grpcHandler := reverse_proxy.NewGrpcLoadBalanceHandler(rb)
			s := grpc.NewServer(
				grpc.ChainStreamInterceptor(
					grpc_proxy_middleware.GrpcServerFlowCountMiddleware(service),
					grpc_proxy_middleware.GrpcClientFlowLimitMiddleware(service),
					grpc_proxy_middleware.GrpcServerFlowLimitMiddleware(service),
					grpc_proxy_middleware.GrpcMetaTransferMiddleware(service),
					grpc_proxy_middleware.GrpcWhiteIplistMiddleware(service),
					grpc_proxy_middleware.GrpcBlackIplistMiddleware(service), ),
				grpc.CustomCodec(proxy.Codec()),
				grpc.UnknownServiceHandler(grpcHandler)) //自定义全局回调

			grpcServerList = append(grpcServerList, s)
			log.Printf(" [INFO] grpc_proxy_run %s\n", ":"+addr)
			if err := s.Serve(lis); err != nil {
				log.Fatalf(" [ERROR] grpc_proxy_run %s err:%v\n", ":"+addr, err)
			}
		}(service)
	}
}

func GrpcServerStop() {
	for _, grpcServer := range grpcServerList {
		grpcServer.GracefulStop()
		log.Printf(" [INFO] grpc_proxy_stop stopped\n")
	}
}
