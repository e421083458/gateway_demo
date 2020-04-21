package http_proxy_middleware

import (
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

func HttpJwtAuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpJwtAuthTokenMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)

		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)
		appMatched := false
		if token != "" {
			claim, err := public.JwtDecode(token)
			if err == nil {
				appList := dao.AppHandler.GetAppList()
				for _, app := range appList {
					if app.AppID == claim.Issuer {
						c.Set("app_detail", app)
						appMatched = true
						break
					}
				}
			}
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && !appMatched {
			middleware.ResponseError(c, 1002, errors.New("HttpJwtAuthTokenMiddleware token error"))
			c.Abort()
			return
		}
		c.Next()
	}
}
