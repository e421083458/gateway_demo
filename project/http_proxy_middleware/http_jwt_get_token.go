package http_proxy_middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

func HttpJwtGetTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/get_token" {
			appID, ok := c.GetPostForm("app_id")
			if !ok {
				middleware.ResponseError(c, 1001, errors.New("need post input app_id"))
				c.Abort()
				return
			}
			secret, ok := c.GetPostForm("secret")
			if !ok {
				middleware.ResponseError(c, 1002, errors.New("need post input secret"))
				c.Abort()
				return
			}
			appList := dao.AppHandler.GetAppList()
			fmt.Println("appList", public.Obj2Json(appList))
			for _, app := range appList {
				if appID != app.AppID {
					continue
				}
				if secret != app.Secret {
					middleware.ResponseError(c, 1003, errors.New("error secret"))
					c.Abort()
					return
				}
				claims := &jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 20).Unix(),
					Issuer:    appID,
				}
				jwtToken, err := public.JwtEncode(claims)
				if err != nil {
					middleware.ResponseError(c, 1004, errors.WithMessage(err, "JwtEncode"))
					c.Abort()
					return
				}
				middleware.ResponseSuccess(c, jwtToken)
				c.Abort()
				return
			}
			middleware.ResponseError(c, 1005, errors.New("not matched app_id"))
			c.Abort()
			return
		}
		c.Next()
	}
}
