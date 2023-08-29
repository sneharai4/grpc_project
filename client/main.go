package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/sneharai4/grpc_project/server/calculator"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorServiceClient(conn)

	expression := "2 + 3" // Your arithmetic expression
	req := &calculator.ExpressionRequest{Expression: expression}

	resp, err := client.EvaluateExpression(context.Background(), req)
	if err != nil {
		log.Fatalf("Error during evaluation: %v", err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
