package public

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

var TransporterHandler *Transporter

type TransParam struct {
	Timeout               time.Duration
	ResponseHeaderTimeout time.Duration
	IdleConnTimeout       time.Duration
	MaxIdleConns          int
}

type Transporter struct {
	transMap     map[string]*http.Transport
	transMapLock sync.RWMutex
}

func init() {
	TransporterHandler = NewTransporter()
}

func NewTransporter() *Transporter {
	return &Transporter{
		transMap:     make(map[string]*http.Transport),
		transMapLock: sync.RWMutex{},
	}
}

func (c *Transporter) GetTrans(name string, param *TransParam) (*http.Transport, error) {
	fmt.Println("GetTransporter Name", name)
	c.transMapLock.Lock()
	defer c.transMapLock.Unlock()
	if trans, ok := c.transMap[name]; ok {
		return trans, nil
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: param.Timeout,
		}).DialContext,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:          param.MaxIdleConns,
		IdleConnTimeout:       param.IdleConnTimeout,
		ResponseHeaderTimeout: param.ResponseHeaderTimeout,
	}
	c.transMap[name] = transport
	return transport, nil
}
