package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CacheKey struct {
	num   int64
	depth int
}

var memo = make(map[CacheKey]int64)

func multiplyRock(input int64) int64 {
	return input * 2024
}

func loadFile(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	content := string(data)
	lines := strings.Split(content, "\r\n")
	return lines
}

func splitInt64(num int64) (int64, int64, bool) {
	numDigits := int(math.Log10(float64(num))) + 1
	if numDigits%2 != 0 {
		return 0, 0, false
	}

	halfDigits := numDigits / 2
	power := int64(math.Pow10(halfDigits))

	firstHalf := num / power
	secondHalf := num % power

	return firstHalf, secondHalf, true
}

func processRock(input int64, iteration int) int64 {
	key := CacheKey{input, iteration}
	if result, found := memo[key]; found {
		return result
	}

	if iteration == 75 {
		return 1
	}

	result := int64(0)
	if input == 0 {
		return processRock(1, iteration+1)
	} else {
		left, right, isEven := splitInt64(input)
		if isEven {
			result = processRock(left, iteration+1) + processRock(right, iteration+1)
		} else {
			result = processRock(multiplyRock(input), iteration+1)
		}
	}

	memo[key] = result
	return result
}

func main() {
	rocksStrings := strings.Split(loadFile(os.Args[1])[0], " ")
	rocks := []int64{}
	for _, rock := range rocksStrings {
		num, err := strconv.ParseInt(rock, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		}
		rocks = append(rocks, num)
	}

	rockLen := int64(0)
	for _, rock := range rocks {
		rockLen += processRock(rock, 0)
	}

	fmt.Println("Solution for the second part is:", rockLen)
}
