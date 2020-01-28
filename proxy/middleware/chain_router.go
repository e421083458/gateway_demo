package middleware

import (
	"net/http"
	"net/http/httputil"
)

// 让 ChainHandlerFunc 继承 http.Handler，方便作其他函数的参数
type ChainHandlerFunc func(rw http.ResponseWriter, req *http.Request)

func (f ChainHandlerFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f(rw, req)
}

// 中间件方法类型
type MiddleWareHandlerFunc func(next http.Handler) http.Handler

// 让 WrapHandlerEntity 继承 http.Handler，方便作其他函数的参数
type WrapHandlerEntity struct {
	Handler http.Handler
}

func (w *WrapHandlerEntity) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	w.Handler.ServeHTTP(rw, req)
}

//代理router，用以创建链式结构
type ChainRouter struct {
	*httputil.ReverseProxy
	prev       *ChainRouter
	middleware MiddleWareHandlerFunc //中间支持
}

func NewChainRouter(p *httputil.ReverseProxy) *ChainRouter {
	return &ChainRouter{
		ReverseProxy: p,
	}
}

//创建 middleware 链式结构
func (p *ChainRouter) Use(middlewares ...MiddleWareHandlerFunc) *ChainRouter {
	if len(middlewares) == 0 {
		return p
	}
	router := p
	for _, mw := range middlewares {
		router = router.use(mw)
	}
	return router
}

//单步链式创建
//尾插法创建链表
func (p *ChainRouter) use(mw MiddleWareHandlerFunc) *ChainRouter {
	return &ChainRouter{
		prev:         p,
		ReverseProxy: p.ReverseProxy,
		middleware:   mw,
	}
}

//基于链表构建方法链
func (p *ChainRouter) genChainFunc(handle http.Handler) http.Handler {
	wraphandler := &WrapHandlerEntity{
		Handler: handle,
	}
	chain := handle
	router := p
	for router.prev != nil {
		if router.middleware != nil {
			//一次调用如下：
			//通过调用倒数第一 middleware 的 MiddleWareHandlerFunc(初始 http.Handler) 获取http.Handler 方法
			//通过调用倒数第二 middleware 的 MiddleWareHandlerFunc(上步所得 http.Handler) 获取 http.Handler 方法
			//...
			//通过调用第一 middleware 的 MiddleWareHandlerFunc(上步所得 http.Handler) 获取http.Handler 方法
			//形成方法的嵌套
			//先加入的在外面

			//形成的方法类似下:
			//ssh := func(h http.Handler, rw http.ResponseWriter, req *http.Request) {
			//	//middware 1 header
			//		//middware 2 header
			//			//middware 3 header
			//				h.ServeHTTP(rw, req)
			//			//middware 3 footer
			//		//middware 2 footer
			//	//middware 1 footer
			//}
			chain = router.middleware(wraphandler)
		}
		wraphandler = &WrapHandlerEntity{
			Handler: chain,
		}
		router = router.prev
	}
	chain = &WrapHandlerEntity{
		Handler: chain,
	}
	return chain
}

//外部服务接口
func (p *ChainRouter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//step 1 基于链表构建方法链
	chainHandler := p.genChainFunc(p.ReverseProxy)
	//step 2 调用方法链
	chainHandler.ServeHTTP(rw, req)
}
