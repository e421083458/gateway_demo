package main

import (
	"flag"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/grpc_proxy_router"
	"github.com/e421083458/gateway_demo/project/http_proxy_router"
	"github.com/e421083458/gateway_demo/project/router"
	"github.com/e421083458/gateway_demo/project/tcp_proxy_router"
	"github.com/e421083458/golang_common/lib"
	"os"
	"os/signal"
	"syscall"
)

var endpoint = flag.String("endpoint", "", "dashboard or server")

func main() {
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *endpoint == "dashboard" {
		lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis",})
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		router.HttpServerStop()
	}
	if *endpoint == "server" {
		lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis",})
		defer lib.Destroy()
		dao.ServiceHandler.LoadOnce()
		go func() {
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			http_proxy_router.HttpSSLServerRun()
		}()
		go func() {
			tcp_proxy_router.TcpServerRun()
		}()
		go func() {
			grpc_proxy_router.GrpcServerRun()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpSSLServerStop()
		tcp_proxy_router.TcpServerStop()
		grpc_proxy_router.GrpcServerStop()
	}
}
