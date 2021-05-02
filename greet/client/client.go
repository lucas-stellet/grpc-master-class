package main

import (
	"context"
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

	doUnary(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lucas",
			LastName:  "Stellet",
		},
	}

	res, err := client.Greet(context.Background(), req)

	if err != nil {
		log.Printf("Request to gRPC fails : %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
