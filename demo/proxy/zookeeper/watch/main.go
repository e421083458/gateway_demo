package main

import (
	"fmt"
	"github.com/e421083458/gateway_demo/proxy/zookeeper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var addr = "127.0.0.1:2002"

func main() {
	//获取zk节点列表
	zkManager := zookeeper.NewZkManager([]string{"127.0.0.1:2181"})
	zkManager.GetConnect()
	defer zkManager.Close()

	zlist, err := zkManager.GetServerListByPath("/real_server")
	fmt.Println("server node:")
	fmt.Println(zlist)
	if err != nil {
		log.Println(err)
	}

	//动态监听节点变化
	chanList, chanErr := zkManager.WatchServerListByPath("/real_server")
	go func() {
		for {
			select {
			case changeErr := <-chanErr:
				fmt.Println("changeErr")
				fmt.Println(changeErr)
			case changedList := <-chanList:
				fmt.Println("watch node changed")
				fmt.Println(changedList)
			}
		}
	}()
	//
	////获取节点内容
	//zc, _, err := zkManager.GetPathData("/rs_server_conf")
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("get node data:")
	//fmt.Println(string(zc))
	//
	////动态监听节点内容
	//dataChan, dataErrChan := zkManager.WatchPathData("/rs_server_conf")
	//go func() {
	//	for {
	//		select {
	//		case changeErr := <-dataErrChan:
	//			fmt.Println("changeErr")
	//			fmt.Println(changeErr)
	//		case changedData := <-dataChan:
	//			fmt.Println("WatchGetData changed")
	//			fmt.Println(string(changedData))
	//		}
	//	}
	//}()

	//关闭信号监听
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
