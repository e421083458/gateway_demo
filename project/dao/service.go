package dao

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/load_balance"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http/httptest"
	"strings"
	"sync"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" validate:"required" toml:"info"`
	HttpRule      *HttpRule      `json:"http_rule" validate:""  toml:"http_rule"`
	TcpRule       *TcpRule       `json:"tcp_rule" validate:""  toml:"tcp_rule"`
	GrpcRule      *GrpcRule      `json:"grpc_rule" validate:""  toml:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" validate:"" toml:"load_balance"`
	AccessControl *AccessControl `json:"access_control" toml:"access_control"`
}

func (t *ServiceDetail) GetTcpLoadBalancer() (load_balance.LoadBalance, error) {
	lb := load_balance.LoadBanlanceFactory(load_balance.LbType(t.LoadBalance.RoundType))
	for index, ip := range t.LoadBalance.GetIPListByModel() {
		lb.Add(ip, t.LoadBalance.GetWeightListByModel()[index])
	}
	lbip, _ := lb.Get("")
	fmt.Println("GetTcpLoadBalancer.Get()", lbip)
	return lb, nil
}

func (t *ServiceDetail) GetHttpLoadBalancer() (load_balance.LoadBalance, error) {
	lb := load_balance.LoadBanlanceFactory(load_balance.LbType(t.LoadBalance.RoundType))
	schema := "http"
	if t.HttpRule.NeedHttps == 1 {
		schema = "https"
	}
	for index, ip := range t.LoadBalance.GetIPListByModel() {
		lb.Add(schema+"://"+ip, t.LoadBalance.GetWeightListByModel()[index])
	}
	lbip, _ := lb.Get("")
	fmt.Println("GetHttpLoadBalancer.Get()", lbip)
	return lb, nil
}

var ServiceHandler *ServiceManager

type ServiceManager struct {
	serviceMap     map[string]*ServiceDetail
	serviceSlice   []*ServiceDetail
	serviceMapLock sync.RWMutex
	init           sync.Once
	err            error
}

func init() {
	ServiceHandler = NewServiceManager()
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		serviceMap:     make(map[string]*ServiceDetail),
		serviceSlice:   []*ServiceDetail{},
		serviceMapLock: sync.RWMutex{},
	}
}

func (s *ServiceManager) GetTcpServiceList() []*ServiceDetail {
	serviceList := []*ServiceDetail{}
	for _, info := range s.serviceSlice {
		if info.Info.LoadType == public.LoadTypeTCP {
			serviceList = append(serviceList, info)
		}
	}
	return serviceList
}

func (s *ServiceManager) GetGrpcServiceList() []*ServiceDetail {
	serviceList := []*ServiceDetail{}
	for _, info := range s.serviceSlice {
		if info.Info.LoadType == public.LoadTypeGRPC {
			serviceList = append(serviceList, info)
		}
	}
	return serviceList
}

func (s *ServiceManager) MatchAccessMode(c *gin.Context) (*ServiceDetail, error) {
	host := c.Request.Host
	uri := c.Request.RequestURI
	path := c.Request.URL.Path
	fmt.Println("c.Request.URL.Host", c.Request.URL.Host)
	fmt.Println("c.Request.URL.Port()", c.Request.URL.Port())
	fmt.Println("c.Request.URL", c.Request.URL)
	fmt.Println("host=", host)
	if strings.Index(host, ":") > 0 {
		host = host[0:strings.Index(host, ":")]
	}
	fmt.Println("host remove port =", host)
	fmt.Println("uri=", uri)
	fmt.Println("path=", path)
	for _, info := range s.serviceSlice {
		if info.Info.LoadType != public.LoadTypeHTTP {
			continue
		}
		if info.HttpRule.RuleType == 1 {
			if info.HttpRule.Rule == host {
				return info, nil
			}
		}
		if info.HttpRule.RuleType == 0 {
			if strings.HasPrefix(path, info.HttpRule.Rule) {
				return info, nil
			}
		}
	}
	return nil, errors.New("not found matched access mode")
}

func (s *ServiceManager) LoadOnce() error {
	s.init.Do(func() {
		model := ServiceInfo{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Set("trace", lib.NewTrace())
		params := &dto.ServiceListInput{PageSize: 9999, PageNo: 1}
		list, _, err := model.ServiceList(c, lib.GORMDefaultPool, params)
		if err != nil {
			s.err = err
			return
		}
		s.serviceMapLock.Lock()
		defer s.serviceMapLock.Unlock()
		for _, item := range list {
			detail, err := item.ServiceDetail(c, lib.GORMDefaultPool, &ServiceInfo{
				ID: item.ID,
			})
			if err != nil {
				s.err = err
				return
			}
			s.serviceMap[item.ServiceName] = detail
			s.serviceSlice = append(s.serviceSlice, detail)
		}
		return
	})
	return s.err
}
