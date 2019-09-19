package main

import (
	"context"
	"log"
	"time"

	pb "../protofiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloGrpcClient(conn)

	name := "Sample"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GreetServer(ctx, &pb.HelloRequest{Greeting: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.Reply)
}
