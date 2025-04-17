package main

import (
	"admin-v1/grpc/helloworldpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type helloServer struct {
	helloworldpb.UnimplementedGreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, req *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(s, &helloServer{})
	reflection.Register(s)

	fmt.Println("Server is running at :50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
