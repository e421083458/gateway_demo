package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HttpJwtClientFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("app_detail")
		if !ok {
			c.Next()
			return
		}

		appDetail := tmp.(*dao.App)
		clientIPLimit := appDetail.Qps
		remoteIP := c.ClientIP()
		if clientIPLimit > 0 {
			limiter, err := public.FlowLimiterHandler.GetLimiter(
				public.FlowAPPPrefix+appDetail.AppID+remoteIP,
				float64(clientIPLimit),
				int(clientIPLimit*3))
			if err != nil {
				middleware.ResponseError(c, 1002, errors.WithMessage(err, "HttpJwtClientFlowLimitMiddleware get GetLimiter error"))
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
