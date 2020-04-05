package middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/proxy/public"
)

func RedisFlowCountMiddleWare(counter *public.RedisFlowCountService) func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		counter.Increase()
		fmt.Println("QPS:", counter.QPS)
		fmt.Println("TotalCount:", counter.TotalCount)
		c.Next()
	}
}
