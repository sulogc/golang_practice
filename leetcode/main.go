package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	s "github.com/sulogc/leetcode"
)

type Solution interface {
	Solve(inputFile string) (string, error)
}

var solutions = map[string]func() Solution{
	"p0001": func() Solution { return &s.P0001{} }, 
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <problem_number>")
		fmt.Println("Example: go run main.go p0001")
		os.Exit(1)
	}

	problemNumber := os.Args[1]
	if !strings.HasPrefix(problemNumber, "p") {
		problemNumber = "p" + fmt.Sprintf("%04d", parseProblemNumber(problemNumber))
	}

	solutionFunc, exists := solutions[problemNumber]
	if !exists {
		fmt.Printf("Error: Solution for %s not found\n", problemNumber+".txt")
		os.Exit(1)
	}
	inputFile := filepath.Join("input", problemNumber+".txt")
	solution := solutionFunc()
	result, err := solution.Solve(inputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Problem: %s\n", problemNumber)
	fmt.Printf("Result: %s\n", result)
}

func parseProblemNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Invalid problem number: %s\n", numStr)
		os.Exit(1)
	}
	return num
}