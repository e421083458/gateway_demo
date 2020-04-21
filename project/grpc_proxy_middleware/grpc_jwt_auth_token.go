package grpc_proxy_middleware

import (
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

func GrpcJwtAuthTokenMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, _ := metadata.FromIncomingContext(ss.Context())
		tokens := md.Get("Authorization")
		token := ""
		if len(tokens) > 0 {
			token = strings.Replace(tokens[0], "Bearer ", "", -1)
		}
		appMatched := false
		if token != "" {
			claim, err := public.JwtDecode(token)
			if err == nil {
				appList := dao.AppHandler.GetAppList()
				for _, app := range appList {
					if app.AppID == claim.Issuer {
						md.Set("app_detail", public.Obj2Json(app))
						appMatched = true
						break
					}
				}
			}
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && !appMatched {
			return errors.New("HttpJwtAuthTokenMiddleware token error")
		}
		err := handler(srv, ss)
		if err != nil {
			log.Printf("RPC failed with error %v\n", err)
		}
		return err
	}
}
