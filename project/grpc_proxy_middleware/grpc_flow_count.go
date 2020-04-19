package grpc_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/public"
	"google.golang.org/grpc"
	"log"
)

//流量统计
func GrpcFlowCountStreamInterceptor(counter *public.RedisFlowCountService) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
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