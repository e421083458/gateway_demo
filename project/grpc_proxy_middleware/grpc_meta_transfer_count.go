package grpc_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

//流量统计
func GrpcMetaTransferMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())
		headerTrans := strings.Split(serviceDetail.GrpcRule.HeaderTransfor, ",")
		if serviceDetail.GrpcRule.HeaderTransfor != "" && len(headerTrans) > 0 {
			fmt.Println("metadata before", md)
			for _, trans := range headerTrans {
				infos := strings.Split(trans, " ")
				if infos[0] == "add" || infos[0] == "edit" {
					md.Set(infos[1], infos[2])
				}
				if infos[0] == "del" {
					delete(md, infos[1])
				}
			}
			ss.SetHeader(md)
			fmt.Println("metadata after", md)
		}

		counter, err := public.FlowCounterHandler.GetCounter(public.FlowCountServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			return err
		}
		counter.Increase()
		err = handler(srv, ss)
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}
