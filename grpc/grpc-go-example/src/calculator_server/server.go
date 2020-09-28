package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-example/src/calculatorpb"
	"io"
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

func (s *server) CalculateSum(
	ctx context.Context,
	req *calculatorpb.CalculateSumRequest,
) (*calculatorpb.CalculateSumResponse, error) {
	sum := req.FirstNumber + req.SecondNumber
	res := &calculatorpb.CalculateSumResponse{Sum: sum}
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
			err := stream.Send(&calculatorpb.DecomposePrimeFactorResponse{PrimeFactor: divisor})
			if err != nil {
				return err
			}
			number = number / divisor
		} else {
			divisor++
		}
	}
	return nil
}

func (s *server) CalculateSumEvenNumber(stream calculatorpb.CalculatorService_CalculateSumEvenNumberServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.CalculateSumEvenNumberResponse{Sum: sum})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		number := req.GetNumber()
		if number%2 == 0 {
			sum += number
		}
	}
}

func (s *server) FindMaxNumber(stream calculatorpb.CalculatorService_FindMaxNumberServer) error {
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		number := req.GetNumber()
		if number > max {
			max = number
			sendErr := stream.Send(&calculatorpb.FindMaxNumberResponse{MaxNumber: max})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}
