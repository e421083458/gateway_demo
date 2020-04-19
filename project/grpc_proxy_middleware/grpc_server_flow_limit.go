package grpc_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
)

func GrpcServerFlowLimitMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		p, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("can not get peer")
		}

		ipIndex := strings.Index(p.Addr.String(), ":")
		clientIP := p.Addr.String()[0:ipIndex]
		fmt.Println("peer.Addr()", clientIP)

		serverLimit := serviceDetail.AccessControl.ServiceFlowLimit
		if serverLimit > 0 {
			limiter, err := public.FlowLimiterHandler.GetLimiter(
				public.FlowCountServicePrefix+serviceDetail.Info.ServiceName+clientIP,
				float64(serverLimit),
				int(serverLimit*3))
			if err != nil {
				return err
			}
			if !limiter.Allow() {
				fmt.Println("not allow")
				return errors.New(fmt.Sprintf("server rate limiting %v,%v", limiter.Limit(), limiter.Burst()))
			}
		}
		err := handler(srv, ss)
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}
