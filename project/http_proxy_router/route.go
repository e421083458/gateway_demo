package http_proxy_router

import (
	"github.com/e421083458/gateway_demo/project/http_proxy_middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(http_proxy_middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(http_proxy_middlewares...)
	router.Group("/")
	router.Use(
		http_proxy_middleware.HttpAccessModeMiddleware(),
		http_proxy_middleware.HttpServerFlowCountMiddleware(),
		http_proxy_middleware.HttpServerFlowLimitMiddleware(),
		http_proxy_middleware.HttpClientFlowLimitMiddleware(),		http_proxy_middleware.HttpHeaderTransferMiddleware(),
		http_proxy_middleware.HttpWhiteIplistMiddleware(),
		http_proxy_middleware.HttpBlackIplistMiddleware(),
		http_proxy_middleware.HttpStripUriMiddleware(),
		http_proxy_middleware.HttpUrlRewriteMiddleware(),
		http_proxy_middleware.HttpReverseProxyMiddleware(), )
	return router
}
