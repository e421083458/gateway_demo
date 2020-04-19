package tcp_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"strings"
)

func TcpServerFlowLimitMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		tmp := c.Get("service_detail")
		if tmp == nil {
			c.conn.Write([]byte("TcpServerFlowCountMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)

		clientIPLimit := serviceDetail.AccessControl.ClientIPFlowLimit

		ipIndex := strings.Index(c.conn.RemoteAddr().String(), ":")
		remoteIP := c.conn.RemoteAddr().String()[0:ipIndex]
		fmt.Println("remoteIP", remoteIP)
		if clientIPLimit > 0 {
			limiter, err := public.FlowLimiterHandler.GetLimiter(
				public.FlowCountServicePrefix+serviceDetail.Info.ServiceName+remoteIP,
				float64(clientIPLimit),
				int(clientIPLimit*3))
			if err != nil {
				c.conn.Write([]byte("TcpServerFlowCountMiddleware get GetLimiter error"))
				c.Abort()
				return
			}
			if !limiter.Allow() {
				fmt.Println("not allow")
				c.conn.Write([]byte(fmt.Sprintf("server rate limiting %v,%v", limiter.Limit(), limiter.Burst())))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
