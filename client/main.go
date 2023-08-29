package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	calculator "github.com/sneharai4/grpc_project/server/calculator"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorServiceClient(conn)

	// Read user input for the expression
	fmt.Print("Enter an arithmetic expression (e.g., 2 + 3): ")
	var expression string
	fmt.Scanln(&expression)
	fmt.Print("express is ", expression)

	req := &calculator.ExpressionRequest{
		Expression: expression,
	}
	fmt.Println("req is ", req.Expression)

	resp, err := client.EvaluateExpression(context.Background(), req)
	if err != nil {
		log.Fatalf("Error during evaluation: %v", err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
