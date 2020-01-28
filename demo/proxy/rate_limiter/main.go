package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	l := rate.NewLimiter(1, 2)

	fmt.Println(l.Limit(), l.Burst())

	for i := 0; i < 100; i++ {
		//阻塞等待直到，取到一个token
		//c, _ := context.WithTimeout(context.Background(), time.Second*2)
		//if err := l.Wait(c); err != nil {
		//	fmt.Println("limiter wait err:" + err.Error())
		//}

		//返回需要等待多久才有新的 token,这样就可以等待指定时间执行任务
		//r := l.Reserve()
		//fmt.Println(r.Delay())

		//判断当前是否可以取到token
		if !l.Allow() {
			fmt.Println("limit no allow")
		}
		time.Sleep(200 * time.Millisecond)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	}
}
