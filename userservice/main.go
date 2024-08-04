package main

import (
	"context"
	"fmt"
	"log"
	"microtask/productservice/service1pb"

	"google.golang.org/grpc"
)

func callService1() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := service1pb.NewService1Client(conn)
	// req := &service1pb.HelloRequest{Name: "Service2"}
	res, err := client.SayHello(context.Background(), nil)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Response from Service 1: %s\n", res.Message)
}

func main() {
    callService1()
    // Rest of the service2 main function...
}