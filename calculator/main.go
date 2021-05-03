package main

import (
	"context"
	"log"
	"net"

	"github.com/lucas-stellet/grpc-learn/grpc-master-class/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	initServergRPC()
}

type Server struct{}

func (s Server) Sum(ctx context.Context, in *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	x := in.GetX()
	y := in.GetY()

	result := x + y

	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func (s Server) PrimeNumberDecomposition(in *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	number := in.GetX()
	divisor := int32(2)
	setResponse := func(result int32) *calculatorpb.PrimeNumberDecompositionResponse {
		return &calculatorpb.PrimeNumberDecompositionResponse{
			Result: divisor,
		}
	}
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(setResponse(divisor))
			number = number / divisor
		} else {
			divisor = divisor + 1
		}
	}
	return nil
}

func initServergRPC() {
	log.Println("calculator server running on port 50051")

	lis, _ := net.Listen("tcp", ":50051")

	grpcServer := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
