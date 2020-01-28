package main

import (
	"github.com/e421083458/gateway_demo/proxy/reverse_proxy_http2/public"
	"github.com/e421083458/gateway_demo/proxy/reverse_proxy_http2/testdata"
	"log"
	"net/http"
	"net/url"
)

var addr = "tonybai.com:3002"

func main() {
	rs1 := "https://tonybai.com:50051"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	urls := []*url.URL{url1}
	proxy := public.NewMultipleHostsReverseProxy(urls)
	log.Println("Starting httpserver at " + addr)
	//log.Fatal(http.ListenAndServe(addr, proxy))
	log.Fatal(http.ListenAndServeTLS(addr, testdata.Path("server.crt"), testdata.Path("server.key"), proxy))
}
