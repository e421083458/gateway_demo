package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

func HttpJwtServerFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("app_detail")
		if !ok {
			c.Next()
			return
		}

		appDetail := tmp.(*dao.App)
		counter, err := public.FlowCounterHandler.GetCounter(public.FlowAPPPrefix + appDetail.AppID)
		if err != nil {
			middleware.ResponseError(c, 1003, errors.WithMessage(err, "HttpJwtServerFlowCountMiddleware get GetCounter error"))
			c.Abort()
			return
		}
		counter.Increase()

		dayCount, err := counter.GetHourCount(time.Now())
		fmt.Println("dayCount", dayCount)
		fmt.Println("GetHourCount.err", err)
		if appDetail.Qpd > 0 && dayCount > appDetail.Qpd {
			middleware.ResponseError(c, 1004, errors.New("total daily requests exceeded"))
			c.Abort()
			return
		}
		c.Next()
	}
}
