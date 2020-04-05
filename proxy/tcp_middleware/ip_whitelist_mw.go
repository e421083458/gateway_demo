package tcp_middleware

import "strings"

func IpWhiteListMiddleWare() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		remoteAddr := c.conn.RemoteAddr().String()
		if strings.HasPrefix(remoteAddr, "127.0.0.1") {
			c.Next()
		} else {
			c.Abort()
			c.conn.Write([]byte("ip_whitelist auth invalid"))
			c.conn.Close()
		}
	}
}

