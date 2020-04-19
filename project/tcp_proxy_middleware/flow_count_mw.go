package tcp_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/public"
)

func FlowCountMiddleWare(counter *public.RedisFlowCountService) func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		counter.Increase()
		fmt.Println("QPS:", counter.QPS)
		fmt.Println("TotalCount:", counter.TotalCount)
		c.Next()
	}
}

