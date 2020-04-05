package main

import (
	"github.com/e421083458/gateway_demo/demo/proxy/reverse_proxy_https/public"
	"github.com/e421083458/gateway_demo/demo/proxy/reverse_proxy_https/testdata"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"net/url"
	"time"
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

	mux := http.NewServeMux()
	mux.Handle("/", proxy)
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 3,            //设置3秒的写超时
		Handler:      mux,
	}
	http2.ConfigureServer(server, &http2.Server{})
	log.Fatal(server.ListenAndServeTLS(testdata.Path("server.crt"), testdata.Path("server.key")))
	log.Fatal(server.ListenAndServe())
}