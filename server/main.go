package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/sneharai4/grpc_project/server/calculator"
)

type server struct{}

func (s *server) EvaluateExpression(ctx context.Context, req *calculator.ExpressionRequest) (*calculator.ResultResponse, error) {
	// Evaluate expression using goroutines and channels
	// For simplicity, let's assume expressions are in the format "2 + 3"
	// You can implement a more complex expression evaluator here

	expression := req.Expression
	var wg sync.WaitGroup
	resultCh := make(chan int32)

	// Split the expression into operands and operator
	// Calculate result using goroutine
	// Send the result to the channel
	// Close the channel after calculation is done
	// Wait for all goroutines to finish

	close(resultCh)

	// Return the result in the response
	return &calculator.ResultResponse{Result: <-resultCh}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(srv, &server{})

	fmt.Println("Server is listening on port 50051")
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
