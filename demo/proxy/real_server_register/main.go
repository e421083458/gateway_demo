package main

import (
	"fmt"
	"github.com/e421083458/gateway_demo/proxy/zookeeper"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rs1 := &RealServer{Addr: "127.0.0.1:2003"}
	rs1.Run()
	time.Sleep(2 * time.Second)
	//rs2 := &RealServer{Addr: "127.0.0.1:2004"}
	//rs2.Run()
	time.Sleep(2 * time.Second)

	//监听关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	log.Println("Starting httpserver at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/error", r.ErrorHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func() {
		//注册zk节点
		zkManager := zookeeper.NewZkManager([]string{"127.0.0.1:2181"})
		err := zkManager.GetConnect()
		if err != nil {
			fmt.Printf(" connect zk error: %s ", err)
		}
		defer zkManager.Close()
		err = zkManager.RegistServerPath("/real_server", r.Addr)
		if err != nil {
			fmt.Printf(" regist node error: %s ", err)
		}
		zlist, err := zkManager.GetServerListByPath("/real_server")
		fmt.Println(zlist)
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	upath := fmt.Sprintf("http://%s%s\n", r.Addr, req.URL.Path)
	io.WriteString(w, upath)
}

func (r *RealServer) ErrorHandler(w http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	w.WriteHeader(500)
	io.WriteString(w, upath)
}
