package main

import (
	"github.com/e421083458/gateway_demo/demo/proxy/reverse_proxy_https/public"
	"github.com/e421083458/gateway_demo/demo/proxy/reverse_proxy_https/testdata"
	"log"
	"net/http"
	"net/url"
)

var addr = "example1.com:3002"

func main() {
	rs1 := "https://example1.com:3003"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	urls := []*url.URL{url1}
	proxy := public.NewMultipleHostsReverseProxy(urls)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServeTLS(addr, testdata.Path("server.crt"), testdata.Path("server.key"), proxy))
}