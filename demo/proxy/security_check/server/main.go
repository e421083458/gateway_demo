package main

import (
	"fmt"
	"github.com/e421083458/gateway_demo/proxy/middleware"
	"github.com/e421083458/gateway_demo/proxy/proxy"
	"github.com/e421083458/gateway_demo/proxy/public"
	"log"
	"net/http"
	"net/url"
)

var addr = "127.0.0.1:2002"

func main() {
	coreFunc := func(c *middleware.SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	sliceRouter := middleware.NewSliceRouter()
	//sliceRouter.Group("/").Use(middleware.IpWhiteListMiddleWare(),)
	sliceRouter.Group("/").Use(middleware.JwtMiddleWare())
	sliceRouter.Group("/get_token").Use(func(c *middleware.SliceRouterContext) {
		user := ""
		secret := ""
		c.Req.ParseForm()
		if len(c.Req.Form["user"]) > 0 {
			user = c.Req.Form["user"][0]
		}
		if len(c.Req.Form["secret"]) > 0 {
			secret = c.Req.Form["secret"][0]
		}
		fmt.Println("user", user)
		fmt.Println("secret", secret)
		if user == "test" && secret == "123abc" {
			jwtToken, err := public.Encode(user)
			if err != nil {
				c.Rw.Write([]byte("get token error:" + err.Error()))
				c.Abort()
				return
			}
			c.Rw.Write([]byte(jwtToken))
			c.Abort()
			return
		}
		c.Rw.Write([]byte("get token error:wrong user or secret"))
		c.Abort()
		return
	})
	routerHandler := middleware.NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}
