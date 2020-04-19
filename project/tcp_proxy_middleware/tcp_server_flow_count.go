package tcp_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
)

func TcpServerFlowCountMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		tmp := c.Get("service_detail")
		fmt.Println("tmp", tmp)
		if tmp == nil {
			c.conn.Write([]byte("TcpServerFlowCountMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)

		counter, err := public.FlowCounterHandler.GetCounter(public.FlowCountServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			c.conn.Write([]byte("HttpServerFlowCountMiddleware get GetCounter error"))
			c.Abort()
			return
		}
		counter.Increase()
		c.Next()
	}
}
