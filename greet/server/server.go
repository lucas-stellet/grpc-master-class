package server

import (
	"context"
	"github.com/lucas-stellet/grpc-master-class/greet/greetpb"
	"fmt"
)

type Server struct {
}

func (s *Server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error)  {
	firstName := in.Greeting.FirstName
	lastName := in.Greeting.LastName

	return &greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello %s %s", firstName, lastName),
	}, nil
}