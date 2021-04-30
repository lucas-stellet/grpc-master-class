package main

import (
	"github.com/lucas-stellet/grpc-master-class/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect :: %v", err)
	}

	defer clientConn.Close()

	gRPCClient := greetpb.NewGreetServiceClient(clientConn)

	log.Println("created client")
}
