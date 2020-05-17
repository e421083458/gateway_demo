package main

import (
	"fmt"
	"github.com/e421083458/gateway_demo/proxy/load_balance"
	proxy2 "github.com/e421083458/gateway_demo/proxy/proxy"
	"github.com/e421083458/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:50055", "40")

	grpcHandler := proxy2.NewGrpcLoadBalanceHandler(rb)
	s := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(grpcHandler)) //自定义全局回调

	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}