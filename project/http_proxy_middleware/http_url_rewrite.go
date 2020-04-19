package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func HttpUrlRewriteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp, ok := c.Get("service_detail")
		if !ok {
			middleware.ResponseError(c, 1001, errors.New("HttpUrlRewriteMiddleware get service_detail error"))
			c.Abort()
			return
		}

		serviceDetail := tmp.(*dao.ServiceDetail)
		fmt.Println("before rewrite", c.Request.URL.Path)
		rewriteRule := strings.Split(serviceDetail.HttpRule.UrlRewrite, ",")
		if len(rewriteRule) > 0 && serviceDetail.HttpRule.UrlRewrite != "" {
			for _, rewrite := range rewriteRule {
				rewriteRule := strings.Split(rewrite, " ")
				regexp, err := regexp.Compile(rewriteRule[0])
				if err != nil {
					middleware.ResponseError(c, 1002, errors.WithMessage(err, "HttpUrlRewriteMiddleware rewriteRules  error"))
					c.Abort()
					return
				}
				rep := regexp.ReplaceAllString(c.Request.URL.Path, rewriteRule[1])
				c.Request.URL.Path = rep
			}
		}
		fmt.Println("after rewrite", c.Request.URL.Path)
		c.Next()
	}
}
