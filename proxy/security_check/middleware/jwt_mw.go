package middleware

import "github.com/e421083458/gateway_demo/proxy/security_check/jwt"

func JwtMiddleWare() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		token := c.Req.Header.Get("token")
		if foo, err := jwt.Decode(token); err == nil && foo == "foo" {
			c.Next()
		} else {
			c.Rw.Write([]byte("jwt auth invalid"))
			c.Abort()
		}
	}
}
