package main

import (
	"fmt"
)

type ModuleConf struct {
	observers []Observer
	conf      []string
	name      string
}

func (s *ModuleConf) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ModuleConf) NotifyAllObservers() {
	for _, obs := range s.observers {
		obs.Update()
	}
}

func (s *ModuleConf) GetConf() []string {
	return s.conf
}

//更新配置时，通知监听者也更新
func (s *ModuleConf) UpdateConf(conf []string) {
	s.conf = conf
	for _, obs := range s.observers {
		obs.Update()
	}
}

func NewModuleConf(name string) (*ModuleConf, error) {
	mConf := &ModuleConf{name: name,}
	return mConf, nil
}

type Observer interface {
	Update()
}

type LoadBalanceObserver struct {
	ModuleConf *ModuleConf
}

func (l *LoadBalanceObserver) Update() {
	fmt.Println("Update get conf:", l.ModuleConf.GetConf())
}

func NewLoadBalanceObserver(ModuleConf *ModuleConf) *LoadBalanceObserver {
	return &LoadBalanceObserver{
		ModuleConf: ModuleConf,
	}
}
