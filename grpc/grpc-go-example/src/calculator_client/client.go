package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-example/src/calculatorpb"
	"log"
)

const (
	host = "localhost"
	port = 50051
)

var (
	client calculatorpb.CalculatorServiceClient
)

func main() {
	connection, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	defer connection.Close()

	client = calculatorpb.NewCalculatorServiceClient(connection)

	calculateSum(10, 20)
}

func calculateSum(firstNumber int32, secondNumber int32) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}
