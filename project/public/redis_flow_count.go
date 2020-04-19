package public

import (
	"context"
	"fmt"
	"github.com/e421083458/golang_common/lib"
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
			tickerCount := atomic.LoadInt64(&reqCounter.TickerCount)
			atomic.StoreInt64(&reqCounter.TickerCount, 0)

			appTodayKey := reqCounter.GetDayKey(time.Now())
			appHourKey := reqCounter.GetHourKey(time.Now())
			if err := RedisConfPipline("default", func(c redis.Conn) {
				c.Send("INCRBY", appTodayKey, tickerCount)
				c.Send("EXPIRE", appTodayKey, 86400*2)
				c.Send("INCRBY", appHourKey, tickerCount)
				c.Send("EXPIRE", appHourKey, 86400*2)
			}); err != nil {
				panic(err)
			}
			totalCount, err := redis.Int64(lib.RedisConfDo(lib.NewTrace(),"default","GET", appTodayKey))
			if err != nil {
				continue
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

func (o *RedisFlowCountService) GetHourKey(t time.Time) string {
	todayHour := t.In(lib.TimeLocation).Format("2006010215")
	return fmt.Sprintf("%s_%s_%s", RedisFlowCountHourKey, o.AppID, todayHour)
}

func (o *RedisFlowCountService) GetDayKey(t time.Time) string {
	today := t.In(lib.TimeLocation).Format("20060102")
	return fmt.Sprintf("%s_%s_%s", RedisFlowCountDayKey, o.AppID, today)
}

func (o *RedisFlowCountService) GetHourCount(t time.Time) (int64, error) {
	return redis.Int64(lib.RedisConfDo(GetTraceContext(context.Background()), "default", "GET", o.GetHourKey(t)))
}

func (o *RedisFlowCountService) GetDayCount(t time.Time) (int64, error) {
	return redis.Int64(lib.RedisConfDo(GetTraceContext(context.Background()), "default", "GET", o.GetDayKey(t)))
}

func (o *RedisFlowCountService) GetQPS() int64 {
	return o.QPS
}

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
