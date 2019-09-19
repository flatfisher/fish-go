package main

import (
	"context"
	"fmt"
	"net"

	"log"

	pb "../protofiles"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GreetServer(ctx context.Context, p *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Request from: %s", p.Greeting)
	return &pb.HelloResponse{Reply: fmt.Sprintf("Hello, %s. ", p.Greeting)}, nil
}

func main() {
	port := 10000
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Run server port: %d", port)
	grpcServer := grpc.NewServer()
	pb.RegisterHelloGrpcServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
