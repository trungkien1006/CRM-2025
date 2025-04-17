package main

import (
	"context"
	"kafka_server/grpc/helloworldpb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := helloworldpb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &helloworldpb.HelloRequest{Name: "alo"})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User: %v", res)
}
