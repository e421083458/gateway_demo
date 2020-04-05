package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"testing"
	"time"
)

func Test_RateLimiter(t *testing.T) {
	l := rate.NewLimiter(1, 5)
	log.Println(l.Limit(), l.Burst())
	for i := 0; i < 100; i++ {
		//阻塞等待直到，取到一个token
		log.Println("before Wait")
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait err:" + err.Error())
		}
		log.Println("after Wait")

		//返回需要等待多久才有新的token,这样就可以等待指定时间执行任务
		r := l.Reserve()
		log.Println("reserve Delay:", r.Delay())

		//判断当前是否可以取到token
		a := l.Allow()
		log.Println("Allow:", a)
	}
}
