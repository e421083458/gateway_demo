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

func GrpcBlackIplistMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		p, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("can not get peer")
		}
		blackList := strings.Split(serviceDetail.AccessControl.BlackList, ",")
		ipIndex := strings.Index(p.Addr.String(), ":")
		clientIP := p.Addr.String()[0:ipIndex]
		fmt.Println("peer.Addr()", clientIP)
		if serviceDetail.AccessControl.OpenAuth == 1 &&
			len(blackList) > 0 &&
			serviceDetail.AccessControl.WhiteList == "" &&
			serviceDetail.AccessControl.BlackList != "" {
			if public.InStringList(clientIP, blackList) {
				return errors.New(clientIP + " in black ip list")
			}
		}
		err := handler(srv, ss)
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}
