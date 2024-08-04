package main

import (
	"fmt"
	"log"
	"microtask/productservice/api"
	"microtask/productservice/service1pb"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service1pb.RegisterService1Server(s, &api.Server{})

	fmt.Println("Service 1 running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
