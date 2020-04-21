package grpc_proxy_middleware

import (
	"encoding/json"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func GrpcJwtServerFlowCountMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())
		tmps := md.Get("app_detail")
		appDetail := &dao.App{}
		if len(tmps) > 0 {
			if err := json.Unmarshal([]byte(tmps[0]), appDetail); err != nil {
				err := handler(srv, ss)
				if err != nil {
					log.Printf("RPC failed with error %v\n", err)
				}
				return err
			}
		}
		counter, err := public.FlowCounterHandler.GetCounter(public.FlowAPPPrefix + appDetail.AppID)
		if err != nil {
			return err
		}
		counter.Increase()

		dayCount, err := counter.GetDayCount(time.Now())
		if appDetail.Qpd > 0 && dayCount > appDetail.Qpd {
			err = errors.New("total daily requests exceeded")
			if err != nil {
				log.Printf("RPC failed with error %v\n", err)
			}
			return err
		}

		err = handler(srv, ss)
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}
