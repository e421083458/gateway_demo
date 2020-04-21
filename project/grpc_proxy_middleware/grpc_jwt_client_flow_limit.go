package grpc_proxy_middleware

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
)

func GrpcJwtClientFlowLimitMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())
		tmps := md.Get("app_detail")
		appDetail:=&dao.App{}
		if len(tmps) > 0 {
			if err := json.Unmarshal([]byte(tmps[0]), appDetail); err != nil {
				err := handler(srv, ss)
				if err != nil {
					log.Printf("RPC failed with error %v\n", err)
				}
				return err
			}
		}

		//https://godoc.org/google.golang.org/grpc/peer#Peer
		p, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("can not get peer")
		}

		ipIndex := strings.Index(p.Addr.String(), ":")
		clientIP := p.Addr.String()[0:ipIndex]
		fmt.Println("peer.Addr()", clientIP)

		clientIPLimit := appDetail.Qps
		if clientIPLimit > 0 {
			limiter, err := public.FlowLimiterHandler.GetLimiter(
				public.FlowAPPPrefix+appDetail.AppID+clientIP,
				float64(clientIPLimit),
				int(clientIPLimit*3))
			if err != nil {
				return err
			}
			if !limiter.Allow() {
				fmt.Println("not allow")
				return errors.New(fmt.Sprintf("client rate limiting %v,%v", limiter.Limit(), limiter.Burst()))
			}
		}
		err := handler(srv, ss)
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}
