package main

import (
	"fmt"
	"log"

	"github.com/lucas-stellet/grpc-learn/grpc-master-class/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("connection has failed when tried to connect: %v", err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	fmt.Printf("Created client: %f", client)
}
