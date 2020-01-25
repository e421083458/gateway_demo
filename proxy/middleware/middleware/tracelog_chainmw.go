package middleware

import (
	"fmt"
	"net/http"
)

func TraceLogChainMW() MiddleWareHandlerFunc {
	return func(next http.Handler) http.Handler {
		return ChainHandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			fmt.Println("trace_in")
			next.ServeHTTP(rw, req)
			fmt.Println("trace_out")
		})
		return nil
	}
}