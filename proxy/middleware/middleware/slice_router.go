package middleware

import (
	"context"
	"math"
	"net/http"
)

//目标定位是 grpc、rpc、http通用的中间件

const abortIndex int8 = math.MaxInt8 / 2 //最多63个中间件

type HandlerFunc func(*SliceRouterContext)

// router结构体
type SliceRouter struct {
	handlers []HandlerFunc
}

// router上下文
type SliceRouterContext struct {
	Rw  http.ResponseWriter
	Req *http.Request
	Ctx context.Context
	*SliceRouter
	index int8
}

func newSliceRouterContext(rw http.ResponseWriter, req *http.Request, r *SliceRouter) *SliceRouterContext {
	newSliceRouter := &SliceRouter{}
	*newSliceRouter = *r //浅拷贝数组指针
	c := &SliceRouterContext{Rw: rw, Req: req, SliceRouter: newSliceRouter, Ctx: req.Context()}
	c.Reset()
	return c
}

func (c *SliceRouterContext) Get(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func (c *SliceRouterContext) Set(key, val interface{}) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

type SliceRouterHandler struct {
	coreFunc func(*SliceRouterContext) http.Handler
	router   *SliceRouter
}

func (w *SliceRouterHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newSliceRouterContext(rw, req, w.router)
	c.handlers = append(c.handlers, func(c *SliceRouterContext) {
		w.coreFunc(c).ServeHTTP(rw, req)
	})
	c.Reset()
	c.Next()
}

func NewSliceRouterHandler(coreFunc func(*SliceRouterContext) http.Handler, router *SliceRouter) *SliceRouterHandler {
	return &SliceRouterHandler{
		coreFunc: coreFunc,
		router:   router,
	}
}

// 构造 router
func NewSliceRouter() *SliceRouter {
	return &SliceRouter{}
}

// 构造回调方法
func (g *SliceRouter) Use(middlewares ...HandlerFunc) *SliceRouter {
	g.handlers = append(g.handlers, middlewares...)
	return g
}

// 从最先加入中间件开始回调
func (c *SliceRouterContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		//fmt.Println("c.index")
		//fmt.Println(c.index)
		c.handlers[c.index](c)
		c.index++
	}
}

// 跳出中间件方法
func (c *SliceRouterContext) Abort() {
	c.index = abortIndex
}

// 是否跳过了回调
func (c *SliceRouterContext) IsAborted() bool {
	return c.index >= abortIndex
}

// 重置回调
func (c *SliceRouterContext) Reset() {
	c.index = -1
}