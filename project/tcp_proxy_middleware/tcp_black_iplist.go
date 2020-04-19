package tcp_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/public"
	"strings"
)

func TcpBlackIplistMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		tmp := c.Get("service_detail")
		if tmp == nil {
			c.conn.Write([]byte("HttpBlackIplistMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)
		blackList := strings.Split(serviceDetail.AccessControl.BlackList, ",")
		ipIndex := strings.Index(c.conn.RemoteAddr().String(), ":")
		clientIP := c.conn.RemoteAddr().String()[0:ipIndex]
		fmt.Println("c.conn.RemoteAddr()", clientIP)
		if serviceDetail.AccessControl.OpenAuth == 1 &&
			len(blackList) > 0 &&
			serviceDetail.AccessControl.WhiteList == "" &&
			serviceDetail.AccessControl.BlackList != "" {
			if public.InStringList(clientIP, blackList) {
				c.conn.Write([]byte(clientIP + " in black ip list"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
