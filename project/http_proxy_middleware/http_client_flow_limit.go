package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HttpClientFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpClientFlowLimitMiddleware get service_detail error"))
			c.Abort()
			return
		}

		serviceDetail := tmp.(*dao.ServiceDetail)
		clientIPLimit := serviceDetail.AccessControl.ClientIPFlowLimit
		remoteIP := c.ClientIP()
		if clientIPLimit > 0 {
			limiter, err := public.FlowLimiterHandler.GetLimiter(
				public.FlowCountServicePrefix+serviceDetail.Info.ServiceName+remoteIP,
				float64(clientIPLimit),
				int(clientIPLimit*3))
			if err != nil {
				middleware.ResponseError(c, 1002, errors.WithMessage(err, "HttpClientFlowLimitMiddleware get GetLimiter error"))
				c.Abort()
				return
			}
			if !limiter.Allow() {
				fmt.Println("not allow")
				middleware.ResponseError(c, 1003, errors.New(fmt.Sprintf("client rate limiting %v,%v", limiter.Limit(), limiter.Burst())))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
