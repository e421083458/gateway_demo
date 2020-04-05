package grpc_interceptor

import (
	"context"
	"fmt"
	"github.com/e421083458/gateway_demo/proxy/public"
	"google.golang.org/grpc"
	"log"
	"time"
)

//流量统计
func GrpcFlowCountUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	counter, _ := public.NewFlowCountService("local_app", time.Second)
	counter.Increase()
	fmt.Println("QPS:", counter.QPS)
	fmt.Println("TotalCount:", counter.TotalCount)
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error %v\n", err)
	}
	return m, err
}

//流量统计
func GrpcFlowCountStreamInterceptor(counter *public.FlowCountService) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		counter.Increase()
		fmt.Println("Grpc Stream QPS:", counter.QPS)
		fmt.Println("Grpc Stream TotalCount:", counter.TotalCount)
		err := handler(srv, newWrappedStream(ss))
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}

