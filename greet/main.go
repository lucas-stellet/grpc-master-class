package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/lucas-stellet/grpc-master-class/greet/greetpb"
	"github.com/lucas-stellet/grpc-master-class/greet/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen :: %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &server.Server{})

	reflection.Register(grpcServer)

	log.Println("gRPC server listen on 50051 ")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("faied to serve gRPC on 50051 :: %v", err)
	}
}
