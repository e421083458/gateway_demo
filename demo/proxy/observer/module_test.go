package main

import (
	"testing"
)

func TestNewLoadBalanceObserver(t *testing.T) {
	moduleConf, err := NewModuleConf("rs_server")
	if err != nil {
		panic(err)
	}
	loadBalanceObserver := NewLoadBalanceObserver(moduleConf)
	moduleConf.Attach(loadBalanceObserver)
	moduleConf.UpdateConf([]string{"122.11.11"})
}
