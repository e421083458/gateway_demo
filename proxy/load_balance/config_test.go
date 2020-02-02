package load_balance

import (
	"testing"
)

func TestNewLoadBalanceObserver(t *testing.T) {
	moduleConf, err := NewLoadBalanceConf("rs_server",
		"/rs_server",
		[]string{"127.0.0.1:2182"},
		map[string]string{"127.0.0.1:2003": "20"})
	if err != nil {
		panic(err)
	}
	loadBalanceObserver := NewLoadBalanceObserver(moduleConf)
	moduleConf.Attach(loadBalanceObserver)
	moduleConf.UpdateConf([]string{"122.11.11"})
}
