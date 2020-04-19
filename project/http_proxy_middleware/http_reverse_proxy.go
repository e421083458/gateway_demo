package http_proxy_middleware

import (
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/gateway_demo/project/reverse_proxy"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

func HttpReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpReverseProxyMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)

		//设置负载均衡策略
		lb, err := serviceDetail.GetHttpLoadBalancer()
		if err != nil {
			middleware.ResponseError(c, 1002, errors.WithMessage(err, "GetHttpLoadBalancer"))
			c.Abort()
			return
		}

		//设置连接池
		transParam := &public.TransParam{
			Timeout:               time.Second * time.Duration(serviceDetail.LoadBalance.UpstreamConnectTimeout),
			ResponseHeaderTimeout: time.Second * time.Duration(serviceDetail.LoadBalance.UpstreamConnectTimeout),
			IdleConnTimeout:       time.Second * time.Duration(serviceDetail.LoadBalance.UpstreamConnectTimeout),
			MaxIdleConns:          serviceDetail.LoadBalance.UpstreamMaxIdle,
		}
		trans, err := public.TransporterHandler.GetTrans(serviceDetail.Info.ServiceName, transParam)
		if err != nil {
			middleware.ResponseError(c, 1003, errors.WithMessage(err, "GetTrans"))
			c.Abort()
			return
		}
		proxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
