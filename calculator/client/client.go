package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/lucas-stellet/grpc-learn/grpc-master-class/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	clientgRPC, gRPCConn := getClientgRPC()
	defer gRPCConn.Close()

	// fmt.Println(CalculatorSumMethod(10, 20, clientgRPC))
	log.Println(CalculatorPrimeNumberDecompositionMethod(120, clientgRPC))
}

func getClientgRPC() (calculatorpb.CalculatorServiceClient, *grpc.ClientConn) {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())

	client := calculatorpb.NewCalculatorServiceClient(conn)

	return client, conn
}

func CalculatorSumMethod(x, y int32, client calculatorpb.CalculatorServiceClient) string {
	response, err := client.Sum(context.Background(), &calculatorpb.SumRequest{X: 10, Y: 20})

	if err != nil {
		log.Println(err)
	}

	return response.String()
}

func CalculatorPrimeNumberDecompositionMethod(x int32, client calculatorpb.CalculatorServiceClient) string {
	var primeDecomposition []string
	stream, err := client.PrimeNumberDecomposition(context.Background(), &calculatorpb.PrimeNumberDecompositionRequest{X: x})

	if err != nil {
		log.Println(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("PrimeNumberDecomposition - error while reading stream: %v", err)
		}

		primeDecomposition = append(primeDecomposition, strconv.FormatInt(int64(res.Result), 10))
	}

	return fmt.Sprintf("Decomposition of prime numbers of %d: %s", x, strings.Join(primeDecomposition, ", "))
}
