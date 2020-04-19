package public

import (
	"sync"
	"time"
)

var FlowCounterHandler *FlowCounter

type FlowCounter struct {
	counterMap     map[string]*RedisFlowCountService
	counterMapLock sync.RWMutex
}

func init() {
	FlowCounterHandler = NewFlowCounter()
}

func NewFlowCounter() *FlowCounter {
	return &FlowCounter{
		counterMap:     make(map[string]*RedisFlowCountService),
		counterMapLock: sync.RWMutex{},
	}
}

func (c *FlowCounter) GetCounter(name string) (*RedisFlowCountService, error) {
	c.counterMapLock.Lock()
	defer c.counterMapLock.Unlock()
	if counter, ok := c.counterMap[name]; ok {
		return counter, nil
	}
	newCounter, err := NewRedisFlowCountService(name, 1*time.Second)
	if err != nil {
		return nil, err
	}
	c.counterMap[name] = newCounter
	return newCounter, nil
}
