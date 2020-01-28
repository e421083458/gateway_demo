package main

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
		Timeout:                1000, // 单次请求 超时时间
		MaxConcurrentRequests:  1,    // 最大并发量
		SleepWindow:            5000, // 熔断后多久去尝试服务是否可用
		RequestVolumeThreshold: 1,    // 验证熔断的 请求数量, 10秒内采样
		ErrorPercentThreshold:  1,    // 验证熔断的 错误百分比
	})

	for i := 0; i < 100; i++ {
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
}
