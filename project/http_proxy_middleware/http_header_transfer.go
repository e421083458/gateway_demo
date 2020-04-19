package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

func HttpHeaderTransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpReverseProxyMiddleware get service_detail error"))
			c.Abort()
			return
		}

		serviceDetail := tmp.(*dao.ServiceDetail)
		headerTrans := strings.Split(serviceDetail.HttpRule.HeaderTransfor, ",")
		fmt.Println("c.header before", c.Request.Header)
		for _, trans := range headerTrans {
			infos := strings.Split(trans, " ")
			if infos[0] == "add" || infos[0] == "edit" {
				c.Request.Header.Set(infos[1], infos[2])
			}
			if infos[0] == "del" {
				c.Request.Header.Del(infos[1])
			}
		}
		fmt.Println("c.header after", c.Request.Header)
		c.Next()
	}
}
