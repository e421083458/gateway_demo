package load_balance

import (
	"fmt"
	"testing"
)

func TestNewLoadBalanceObserver(t *testing.T) {
	moduleConf, err := NewLoadBalanceZkConf("%s",
		"/real_server",
		[]string{"127.0.0.1:2181"},
		map[string]string{"127.0.0.1:2003": "20"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	loadBalanceObserver := NewLoadBalanceObserver(moduleConf)
	moduleConf.Attach(loadBalanceObserver)
	moduleConf.UpdateConf([]string{"122.11.11"})
	select {}
}