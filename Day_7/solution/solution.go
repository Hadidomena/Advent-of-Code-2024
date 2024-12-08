package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helper function to evaluate an expression in strict left-to-right order
func evaluateLeftToRight(numbers []string, operators []string) int64 {
	result, err := strconv.ParseInt(numbers[0], 10, 64) // Start with the first number
	if err != nil {
		fmt.Println("Error parsing number:", err)
		return 0
	}

	// Iterate through numbers and operators to compute the result
	for i := 0; i < len(operators); i++ {
		nextNum, err := strconv.ParseInt(strings.TrimSpace(numbers[i+1]), 10, 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return 0
		}
		if operators[i] == "+" {
			result += nextNum
		} else if operators[i] == "*" {
			result *= nextNum
		} else if operators[i] == "||" {
			str := strconv.FormatInt(result, 10)
			joined := str + numbers[i+1]
			number, err := strconv.ParseInt(strings.TrimSpace(joined), 10, 64)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return 0
			}
			result = number
		}
	}

	return result
}

// Helper function to generate all combinations of + and * operators
func generateCombinations(numbers []string, operators []string, idx int, target int64) bool {
	// Base case: if all operators are placed, evaluate the expression
	if idx == len(numbers)-1 {
		return evaluateLeftToRight(numbers, operators) == target
	}

	// Recursive case: try adding "+" and "*" operators
	return generateCombinations(numbers, append(operators, "+"), idx+1, target) ||
		generateCombinations(numbers, append(operators, "*"), idx+1, target) ||
		generateCombinations(numbers, append(operators, "||"), idx+1, target)
}

// Main function to check if a target is achievable
func checkIfPossible(target int64, numbers []string) bool {
	// Start the recursive operator placement
	return generateCombinations(numbers, []string{}, 0, target)
}

func main() {
	// Read the input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the file content
	var result int64
	content := string(data)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, ": ")
		numbers := strings.Split(splitLine[1], " ")
		num, err := strconv.ParseInt(splitLine[0], 10, 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		if checkIfPossible(num, numbers) {
			result += num
		}
	}

	fmt.Println("Result:", result)
}
