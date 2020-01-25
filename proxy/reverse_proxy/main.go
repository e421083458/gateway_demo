package main

import (
	"log"
	"net/http"
	"net/url"
)

var addr = "127.0.0.1:2002"

func main() {
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
	proxy := NewMultipleHostsReverseProxy(urls)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
