package middleware

import (
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AdminSessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get(public.AdminInfoSessionKey) == nil {
			ResponseError(c, 200, errors.New("管理端未登陆"))
			c.Abort()
			return
		}
		c.Next()
	}
}
