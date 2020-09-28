package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-example/src/calculatorpb"
	"log"
	"net"
)

const (
	host = "localhost"
	port = 50051
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Serve failed: %v", err)
	}
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	sum := req.FirstNumber + req.SecondNumber
	res := &calculatorpb.SumResponse{SumResult: sum}
	return res, nil
}

func (s *server) DecomposePrimeFactor(
	req *calculatorpb.DecomposePrimeFactorRequest,
	stream calculatorpb.CalculatorService_DecomposePrimeFactorServer,
) error {
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			if err := stream.Send(&calculatorpb.DecomposePrimeFactorResponse{
				PrimeFactor: divisor,
			}); err != nil {
				return err
			}
			number = number / divisor
		} else {
			divisor++
		}
	}
	return nil
}
