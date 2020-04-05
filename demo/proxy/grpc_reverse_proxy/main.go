package main

import (
	"context"
	"fmt"
	"github.com/e421083458/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"strings"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// 拒绝某些特殊请求
		if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
			return ctx, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			// 基于header决定下游请求
			if val, exists := md[":authority"]; exists && val[0] == "staging.api.example.com" {
				c, err := grpc.DialContext(ctx, "api-service.staging.svc.local", grpc.WithCodec(proxy.Codec()),grpc.WithInsecure())
				return ctx, c, err
			} else {
				c, err := grpc.DialContext(ctx, "localhost:50055", grpc.WithCodec(proxy.Codec()),grpc.WithInsecure())
				return ctx, c, err
			}
		}
		return ctx, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}

	s := grpc.NewServer(grpc.CustomCodec(proxy.Codec()), grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
