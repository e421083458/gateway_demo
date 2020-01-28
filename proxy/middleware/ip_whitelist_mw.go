package middleware

import (
	"net"
)

func IpWhiteListMiddleWare() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		remoteAddr, _, _ := net.SplitHostPort(c.Req.RemoteAddr)
		if remoteAddr == "127.0.0.1" {
			c.Next()
		} else {
			c.Abort()
			c.Rw.Write([]byte("ip_whitelist auth invalid"))
		}
	}
}
