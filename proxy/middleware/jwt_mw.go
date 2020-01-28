package middleware

import "github.com/e421083458/gateway_demo/proxy/public"

func JwtMiddleWare() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		token := c.Req.Header.Get("token")
		if foo, err := public.Decode(token); err == nil && foo == "foo" {
			c.Next()
		} else {
			c.Rw.Write([]byte("jwt auth invalid"))
			c.Abort()
		}
	}
}
