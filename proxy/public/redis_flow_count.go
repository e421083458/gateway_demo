package public

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync/atomic"
	"time"
)

type RedisFlowCountService struct {
	AppID       string
	Interval    time.Duration
	QPS         int64
	Unix        int64
	TickerCount int64
	TotalCount  int64
}

func NewRedisFlowCountService(appID string, interval time.Duration) (*RedisFlowCountService, error) {
	reqCounter := &RedisFlowCountService{
		AppID:    appID,
		Interval: interval,
		QPS:      0,
		Unix:     0,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			tickerCount := atomic.LoadInt64(&reqCounter.TickerCount)     //获取数据
			atomic.StoreInt64(&reqCounter.TickerCount, 0)            //重置数据

			tl, _ := time.LoadLocation("Asia/Chongqing")
			today := time.Now().In(tl).Format("2006-01-02")
			totalAppKey := fmt.Sprintf("%s_%s_%s", "totalcall", today, appID)
			if err := RedisConfPipline(func(c redis.Conn) {
				c.Send("INCRBY", totalAppKey, tickerCount)
				c.Send("EXPIRE", totalAppKey, 86400)
			}); err != nil {
				panic(err)
			}

			totalCount, err := redis.Int64(RedisConfDo("GET", totalAppKey))
			if err != nil {
				panic(err)
			}
			nowUnix := time.Now().Unix()
			if reqCounter.Unix == 0 {
				reqCounter.Unix = time.Now().Unix()
				continue
			}
			tickerCount = totalCount - reqCounter.TotalCount
			if nowUnix > reqCounter.Unix {
				reqCounter.TotalCount = totalCount
				reqCounter.QPS = tickerCount / (nowUnix - reqCounter.Unix)
				reqCounter.Unix = time.Now().Unix()
			}
		}
	}()
	return reqCounter, nil
}

//原子增加
func (o *RedisFlowCountService) Increase() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		atomic.AddInt64(&o.TickerCount, 1)
	}()
}
