package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "127.0.0.1:2002"
	defaultName = "world"
)

func main() {
	//ts, err := credentials.NewClientTLSFromFile("../server.pem", "")
	//if err != nil {
	//	log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	//}
	//conn, err := grpc.Dial(address,grpc.WithTransportCredentials(ts),)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}