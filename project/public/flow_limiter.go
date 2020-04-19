package public

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync"
)

var FlowLimiterHandler *FlowLimiter

type FlowLimiter struct {
	limiterMap     map[string]*rate.Limiter
	limiterMapLock sync.RWMutex
}

func init() {
	FlowLimiterHandler = NewFlowLimiter()
}

func NewFlowLimiter() *FlowLimiter {
	return &FlowLimiter{
		limiterMap:     make(map[string]*rate.Limiter),
		limiterMapLock: sync.RWMutex{},
	}
}

func (c *FlowLimiter) GetLimiter(name string, limit float64, burst int) (*rate.Limiter, error) {
	fmt.Println("GetLimiter Name", name)
	c.limiterMapLock.Lock()
	defer c.limiterMapLock.Unlock()
	if counter, ok := c.limiterMap[name]; ok {
		return counter, nil
	}
	newLimiter := rate.NewLimiter(rate.Limit(limit), burst)
	c.limiterMap[name] = newLimiter
	return newLimiter, nil
}
