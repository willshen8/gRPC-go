package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/willshen8/gRPC-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &calculatorpb.SumRequest{
		A: &calculatorpb.Integer{
			A: 100,
		},
		B: &calculatorpb.Integer{
			A: 200,
		},
	}
	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do server streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Prime: 1246450,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecompositionRequest RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happend: %v", err)
		}
		log.Println(res.GetPrimeFactor())
	}
}
