package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/willshen8/gRPC-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator function was invoked with %v", req)
	a := req.A.GetA()
	b := req.B.GetA()
	var result = a + b
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)
	number := req.GetPrime()
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			log.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func main() {
	fmt.Println("Calculator server has started...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
