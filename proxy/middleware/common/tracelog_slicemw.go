package common

import (
	"log"
)

func TraceLogSliceMW() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		log.Println("trace_in")
		log.Println("Next")
		c.Next()
		log.Println("trace_out")
	}
}