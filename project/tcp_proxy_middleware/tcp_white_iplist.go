package tcp_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"strings"
)

func TcpWhiteIplistMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		tmp := c.Get("service_detail")
		if tmp == nil {
			c.conn.Write([]byte("HttpBlackIplistMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)
		whiteList := strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		ipIndex := strings.Index(c.conn.RemoteAddr().String(), ":")
		clientIP := c.conn.RemoteAddr().String()[0:ipIndex]
		fmt.Println("c.conn.RemoteAddr()", clientIP)
		if serviceDetail.AccessControl.OpenAuth == 1 &&
			len(whiteList) > 0 &&
			serviceDetail.AccessControl.WhiteList != "" {
			if !public.InStringList(clientIP, whiteList) {
				c.conn.Write([]byte(clientIP + " not in white ip list"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
