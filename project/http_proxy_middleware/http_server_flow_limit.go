package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HttpServerFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpReverseProxyMiddleware get service_detail error"))
			c.Abort()
			return
		}

		serviceDetail := tmp.(*dao.ServiceDetail)
		serverLimit := serviceDetail.AccessControl.ServiceFlowLimit
		if serverLimit > 0 {
			limiter, err := public.FlowLimiterHandler.GetLimiter(public.FlowServicePrefix+serviceDetail.Info.ServiceName, float64(serverLimit), int(serverLimit*3))
			if err != nil {
				middleware.ResponseError(c, 1002, errors.WithMessage(err, "HttpServerFlowLimitMiddleware get GetLimiter error"))
				c.Abort()
				return
			}
			if !limiter.Allow() {
				middleware.ResponseError(c, 1003, errors.New(fmt.Sprintf("server rate limiting %v,%v", limiter.Limit(), limiter.Burst())))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
