package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Read user input for the expression
		fmt.Print("Enter an arithmetic expression (e.g., 2 + 3 or exit to EXIT): ")
		scanner.Scan()
		expression := scanner.Text()

		if expression == "exit" {
			break
		}

		// Trim spaces around the expression
		expression = strings.TrimSpace(expression)
		req := &calculator.ExpressionRequest{
			Expression: expression,
		}

		resp, err := client.EvaluateExpression(context.Background(), req)
		if err != nil {
			log.Fatalf("Error during evaluation: %v", err)
		}

		fmt.Printf("Result: %d\n", resp.Result)
	}

}
