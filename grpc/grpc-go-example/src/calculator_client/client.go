package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-example/src/calculatorpb"
	"io"
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
	decomposePrimeFactor(7192285800)
	calculateSumEvenNumber([]int32{2, 3, 4, 9, 10, 18, 23, 30})
}

func calculateSum(firstNumber int32, secondNumber int32) {
	req := &calculatorpb.CalculateSumRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}

	res, err := client.CalculateSum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CalculateSum: %v", err)
	}
	log.Printf("Response from CalculateSum: %v", res.GetSum())
}

func decomposePrimeFactor(number int64) {
	req := &calculatorpb.DecomposePrimeFactorRequest{Number: number}
	stream, err := client.DecomposePrimeFactor(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DecomposePrimeFactor: %v", err)
	}

	result := ""
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving response: %v", err)
		}
		result += fmt.Sprintf("%d ", res.GetPrimeFactor())
	}
	log.Printf("Response from DecomposePrimeFactor: %s", result)
}

func calculateSumEvenNumber(numbers []int32) {
	stream, err := client.CalculateSumEvenNumber(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream CalculateSumEvenNumber: %v", err)
	}

	for _, number := range numbers {
		err = stream.Send(&calculatorpb.CalculateSumEvenNumberRequest{Number: number})
		if err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("Response from CalculateSumEvenNumberRequest: %d", res.GetSum())
}
