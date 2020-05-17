// Binary client is an example client.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/e421083458/gateway_demo/demo/proxy/grpc_server_client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var addr = flag.String("addr", "localhost:8402", "the address to connect to")

const (
	timestampFormat = time.StampNano // "Jan _2 15:04:05.000"
	streamingCount  = 10
	AccessToken="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk2OTExMTQsImlzcyI6ImFwcF9pZF9iIn0.qb2A_WsDP_-jfQBxJk6L57gTnAzZs-SPLMSS_UO6Gkc"
)

func unaryCallWithMetadata(c pb.EchoClient, message string) {
	fmt.Printf("--- unary ---\n")

	// Create metadata and context.
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("failed to call UnaryEcho: %v", err)
	}
	fmt.Printf("response:%v\n", r.Message)
}

func serverStreamingWithMetadata(c pb.EchoClient, message string) {
	fmt.Printf("--- server streaming ---\n")

	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := c.ServerStreamingEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("failed to call ServerStreamingEcho: %v", err)
	}

	// Read all the responses.
	var rpcStatus error
	fmt.Printf("response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Message)
	}
	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server streaming: %v", rpcStatus)
	}
}

func clientStreamWithMetadata(c pb.EchoClient, message string) {
	fmt.Printf("--- client streaming ---\n")
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.ClientStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call ClientStreamingEcho: %v\n", err)
	}

	// Send all requests to the server.
	for i := 0; i < streamingCount; i++ {
		if err := stream.Send(&pb.EchoRequest{Message: message}); err != nil {
			log.Fatalf("failed to send streaming: %v\n", err)
		}
	}

	// Read the response.
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to CloseAndRecv: %v\n", err)
	}
	fmt.Printf("response:%v\n", r.Message)
}

func bidirectionalWithMetadata(c pb.EchoClient, message string) {
	fmt.Printf("--- bidirectional ---\n")
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call BidirectionalStreamingEcho: %v\n", err)
	}

	go func() {
		// Send all requests to the server.
		for i := 0; i < streamingCount; i++ {
			if err := stream.Send(&pb.EchoRequest{Message: message}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// Read all the responses.
	var rpcStatus error
	fmt.Printf("response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Message)
	}
	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server streaming: %v", rpcStatus)
	}
}

const message = "this is examples/metadata"

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := grpc.Dial(*addr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			c := pb.NewEchoClient(conn)

			//调用一元方法
			//for i := 0; i < 100; i++ {
			unaryCallWithMetadata(c, message)
			time.Sleep(400 * time.Millisecond)
			//}
			//
			//服务端流式
			serverStreamingWithMetadata(c, message)
			time.Sleep(1 * time.Second)

			//客户端流式
			clientStreamWithMetadata(c, message)
			time.Sleep(1 * time.Second)

			//双向流式
			bidirectionalWithMetadata(c, message)
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
}
