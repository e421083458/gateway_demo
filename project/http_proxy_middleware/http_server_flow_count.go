package http_proxy_middleware

import (
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HttpServerFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpServerFlowCountMiddleware get service_detail error"))
			c.Abort()
			return
		}

		serviceDetail := tmp.(*dao.ServiceDetail)

		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			middleware.ResponseError(c, 1002, errors.WithMessage(err, "HttpServerFlowCountMiddleware get GetCounter error"))
			c.Abort()
			return
		}
		totalCounter.Increase()

		counter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			middleware.ResponseError(c, 1003, errors.WithMessage(err, "HttpServerFlowCountMiddleware get GetCounter error"))
			c.Abort()
			return
		}

		counter.Increase()
		c.Next()
	}
}
