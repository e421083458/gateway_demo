package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

func HttpWhiteIplistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpWhiteIplistMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)
		whiteList := strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		fmt.Println("c.ClientIP()", c.ClientIP())
		if serviceDetail.AccessControl.OpenAuth==1 &&
			len(whiteList) > 0 &&
			serviceDetail.AccessControl.WhiteList != "" {
			if !public.InStringList(c.ClientIP(), whiteList) {
				middleware.ResponseError(c, 1002, errors.New(c.ClientIP()+" not in white ip list"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
