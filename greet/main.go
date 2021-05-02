package main

import (
	"log"
	"net"

	"github.com/lucas-stellet/grpc-learn/grpc-master-class/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("grpc server started at 50051 port")
}
