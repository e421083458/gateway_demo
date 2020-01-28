package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	//"github.com/idubinskiy/http-to-grpc-gateway/go/hello"
)

const (
	httpPort = 9000
	grpcPort = 9010

	oldServiceAddress = "http://localhost:8000"
)

func main() {
	// create TCP listener for gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}

	// create gRPC server and register service on it
	//grpcServer := grpc.NewServer()
	//hello.RegisterHelloServer(grpcServer, server{})

	// start listening for gRPC requests in a goroutine
	//go func() {
	//	err := grpcServer.Serve(listener)
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	//log.Printf("Go gRPC server listening on port %d!", grpcPort)

	// create grpc-gateway mux
	grpcGatewayMux := runtime.NewServeMux()

	// add source header, log requests
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Source-Server", "go")
		log.Printf("Got request for path: %s, method: %s", r.URL.Path, r.Method)
		grpcGatewayMux.ServeHTTP(w, r)
	})

	// create reverse proxy to Node service
	oldServiceURL, err := url.Parse(oldServiceAddress)
	if err != nil {
		panic(err)
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(oldServiceURL)

	// serve unimplemented route/method combinations using reverse proxy
	runtime.OtherErrorHandler = func(w http.ResponseWriter, r *http.Request, msg string, code int) {
		if code != http.StatusNotFound && code != http.StatusMethodNotAllowed {
			runtime.DefaultOtherErrorHandler(w, r, msg, code)
			return
		}

		log.Printf("Proxying request for path: %s, method: %s", r.URL.Path, r.Method)

		// remove header; Node will add its own
		w.Header().Del("X-Source-Server")
		reverseProxy.ServeHTTP(w, r)
	}

	// register grpc-gateway service on grpc-gateway mux, proxying requests to gRPC server
	//err = hello.RegisterHelloHandlerFromEndpoint(
	//	context.Background(),
	//	grpcGatewayMux,
	//	fmt.Sprintf("localhost:%d", grpcPort),
	//	[]grpc.DialOption{grpc.WithInsecure()},
	//)
	//if err != nil {
	//	panic(err)
	//}

	// start listening for HTTP requests in a goroutine
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
		if err != nil {
			panic(err)
		}
	}()
	log.Printf("Go HTTP server listening on port %d!", httpPort)

	// block forever, letting other goroutines run
	select {}
}
//
//type server struct{}
//
//func (s server) Hello(ctx context.Context, req *empty.Empty) (resp *hello.Response, err error) {
//	return &hello.Response{
//		Message: "Hello World!",
//	}, nil
//}
//
//func (s server) HelloName(ctx context.Context, req *hello.NameRequest) (resp *hello.Response, err error) {
//	return &hello.Response{
//		Message: fmt.Sprintf("Hello %s!", req.Name),
//	}, nil
//}
