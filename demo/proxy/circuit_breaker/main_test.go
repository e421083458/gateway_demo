package main

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8074", hystrixStreamHandler)

	hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
		Timeout:                1000, // 单次请求 超时时间
		MaxConcurrentRequests:  1,    // 最大并发量
		SleepWindow:            5000, // 熔断后多久去尝试服务是否可用
		RequestVolumeThreshold: 1,    // 验证熔断的 请求数量, 10秒内采样
		ErrorPercentThreshold:  1,    // 验证熔断的 错误百分比
	})

	for i := 0; i < 10000; i++ {
		//异步调用使用 hystrix.Go
		err := hystrix.Do("aaa", func() error {
			//test case 1 并发测试
			if i == 0 {
				return errors.New("service error")
			}
			//test case 2 超时测试
			//time.Sleep(2 * time.Second)
			log.Println("do services")
			return nil
		}, nil)
		if err != nil {
			log.Println("hystrix err:" + err.Error())
			time.Sleep(1 * time.Second)
			log.Println("sleep 1 second")
		}
	}
	time.Sleep(100 * time.Second)
}
