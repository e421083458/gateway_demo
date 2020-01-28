package middleware

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
)

func CircuitMW() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		err := hystrix.Do("common", func() error {
			c.Next()
			statusCode, ok := c.Get("status_code").(int)
			if !ok || statusCode != 200 {
				return errors.New("downstream error")
			}
			return nil
		}, nil)
		if err != nil {
			switch err {
			case hystrix.ErrCircuitOpen:
				c.Rw.Write([]byte("circuit error:" + err.Error()))
			case hystrix.ErrMaxConcurrency:
				c.Rw.Write([]byte("circuit error:" + err.Error()))
			default:

			}
			c.Abort()
		}
	}
}
