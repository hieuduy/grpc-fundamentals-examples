package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpc-gateway-go-example/src/calculatorpb"
	"log"
	"net"
	"net/http"
)

const (
	restHost = "localhost"
	restPort = 8080
	grpcHost = "localhost"
	grpcPort = 50051
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", restHost, restPort))
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}

	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = calculatorpb.RegisterCalculatorServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("%s:%d", grpcHost, grpcPort),
		dialOptions,
	)
	if err != nil {
		log.Fatalf("Register service handler failed: %v", err)
	}

	if err := http.Serve(listener, mux); err != nil {
		log.Fatalf("Serve failed: %v", err)
	}
}
