package middleware

import (
	"log"
)

func TraceLogSliceMW() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		log.Println("trace_in")
		c.Next()
		log.Println("trace_out")
	}
}
