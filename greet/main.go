package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lucas-stellet/grpc-learn/grpc-master-class/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("grpc server started at 50051 port")

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type server struct{}

func (s *server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("greet function was invoked with %v\n", in)
	firstName := in.GetGreeting().GetFirstName()
	lastName := in.GetGreeting().GetLastName()

	result := fmt.Sprintf("Hello %s %s", firstName, lastName)

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}
