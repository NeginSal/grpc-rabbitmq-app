package main

import (
	"log"
	"net"

	"github.com/NeginSal/grpc-rabbitmq-app/proto/orderpb"
	"github.com/NeginSal/grpc-rabbitmq-app/order"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	orderpb.RegisterOrderServiceServer(grpcServer, &order.Server{})

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}