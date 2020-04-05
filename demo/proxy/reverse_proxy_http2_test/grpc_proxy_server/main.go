package main

import (
	"context"
	"github.com/e421083458/gateway_demo/proxy/tcp_proxy"
	"google.golang.org/grpc/reverse_proxy"
	"log"
	"math"
	"net"
)

var (
	addr = "127.0.0.1:2002"
)

const (
	defaultServerMaxReceiveMessageSize = 1024 * 1024 * 4
	defaultServerMaxSendMessageSize    = math.MaxInt32
)

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, c net.Conn) {

	////rawConn.SetDeadline(time.Now().Add(s.opts.connectionTimeout))
	////conn, authInfo, err := s.useTransportAuthenticator(rawConn)
	////return rawConn, nil, nil
	////st := s.newHTTP2Transport(conn, authInfo)
	reverse_proxy.ServeTCP(ctx, c)
	//config := &transport.ServerConfig{
	//	MaxStreams:            0,
	//	AuthInfo:              nil,
	//	InTapHandle:           nil,
	//	StatsHandler:          nil,
	//	KeepaliveParams:       keepalive.ServerParameters{},
	//	KeepalivePolicy:       keepalive.EnforcementPolicy{},
	//	InitialWindowSize:     0,
	//	InitialConnWindowSize: 0,
	//	WriteBufferSize:       defaultServerMaxSendMessageSize,
	//	ReadBufferSize:        defaultServerMaxSendMessageSize,
	//	ChannelzParentID:      0,
	//	MaxHeaderListSize:     nil,
	//	HeaderTableSize:       nil,
	//}
	//st, err := transport.NewServerTransport("http2", c, config)
	//if err != nil {
	//	fmt.Println("NewServerTransport(%q) failed: %v", c.RemoteAddr(), err)
	//	//c.Close()
	//	return
	//}
	////rawConn.SetDeadline(time.Time{})
	////s.serveStreams(st)
	//defer st.Close()
	//var wg sync.WaitGroup
	//st.HandleStreams(func(stream *transport.Stream) {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		sm := stream.Method()
	//		if sm != "" && sm[0] == '/' {
	//			sm = sm[1:]
	//		}
	//		pos := strings.LastIndex(sm, "/")
	//		if pos == -1 {
	//			errDesc := fmt.Sprintf("malformed method name: %q", stream.Method())
	//			fmt.Println(errDesc)
	//			return
	//		}
	//		service := sm[:pos]
	//		method := sm[pos+1:]
	//		fmt.Println("method", method)
	//		fmt.Println("service", service)
	//		//s.handleStream(st, stream, s.traceInfo(st, stream))
	//	}()
	//}, func(ctx context.Context, method string) context.Context {
	//	return ctx
	//	//if !EnableTracing {
	//	//	return ctx
	//	//}
	//	//tr := trace.New("grpc.Recv."+methodFamily(method), method)
	//	//return trace.NewContext(ctx, tr)
	//})
	//wg.Wait()
	////if !s.addConn(st) {
	////	return
	////}
	////s.serveStreams(st)
}

func main() {
	tcpHandler:=&tcpHandler{}
	log.Println("Starting tcpserver at " + addr)
	log.Fatal(tcp_proxy.ListenAndServe(addr, tcpHandler))
}
