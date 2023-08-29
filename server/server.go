package main

import (
	"context"
	"fmt"
	calculator "github.com/sneharai4/grpc_project/server/calculator"
	"google.golang.org/grpc"
	"net"
	"regexp"
	"strconv"
)

type server struct {
	calculator.UnimplementedCalculatorServiceServer
}

func evaluateExpression(expression string, resultCh chan<- int32) {
	defer close(resultCh)

	// Use regular expression to capture operands and operator
	regex := regexp.MustCompile(`\s*([-+*/])\s*`)
	matches := regex.FindAllStringSubmatch(expression, -1)
	if len(matches) != 1 || len(matches[0]) != 2 {
		resultCh <- 0 // Invalid expression
		return
	}

	match := matches[0]
	operator := match[1]
	operands := regex.Split(expression, -1)
	operand1, err1 := strconv.Atoi(operands[0])
	operand2, err2 := strconv.Atoi(operands[1])

	if err1 != nil || err2 != nil {
		resultCh <- 0 // Invalid operands
		return
	}

	var result int32

	switch operator {
	case "+":
		result = int32(operand1 + operand2)
	case "-":
		result = int32(operand1 - operand2)
	case "*":
		result = int32(operand1 * operand2)
	case "/":
		if operand2 != 0 {
			result = int32(operand1 / operand2)
		} else {
			resultCh <- 0 // Division by zero
			return
		}
	default:
		resultCh <- 0 // Invalid operator
		return
	}

	resultCh <- result
}

func (s *server) EvaluateExpression(ctx context.Context, req *calculator.ExpressionRequest) (*calculator.ResultResponse, error) {
	resultCh := make(chan int32)

	go evaluateExpression(req.Expression, resultCh)

	result := <-resultCh
	return &calculator.ResultResponse{Result: result}, nil
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
