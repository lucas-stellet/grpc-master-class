package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lucas-stellet/grpc-learn/grpc-master-class/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	clientgRPC, gRPCConn := getClientgRPC()
	defer gRPCConn.Close()

	response, err := clientgRPC.Sum(context.Background(), &calculatorpb.SumRequest{X: 10, Y: 20})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(response.String())
}

func getClientgRPC() (calculatorpb.CalculatorServiceClient, *grpc.ClientConn) {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())

	client := calculatorpb.NewCalculatorServiceClient(conn)

	return client, conn
}
