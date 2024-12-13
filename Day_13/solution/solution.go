package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Hadidomena/AoC_Go_utils/utils"
)

func splitLine(input string, isIncreased bool) [2]int64 {
	buttons := strings.Split(strings.Split(input, ": ")[1], ", ")
	result := [2]int64{}
	for i, button := range buttons {
		num, err := strconv.ParseInt(button[2:], 10, 64)
		if err == nil {
			if isIncreased {
				result[i] = num + 10000000000000
			} else {
				result[i] = num
			}
		}
	}
	return result
}
func FindCombination(vec1, vec2, target [2]int64) (int64, int64) {
	// I calculated a and b by solving equation a * vec1 + b * vec2 == target
	// where you can see that a * vec1[0] + b * vec2[0] == target[0]
	// and a * vec1[1] + b * vec2[1] == target[1]
	a := float64(float64(target[0]*vec2[1]-target[1]*vec2[0]) / float64(vec1[0]*vec2[1]-vec1[1]*vec2[0]))
	b := float64(float64(target[1])-a*float64(vec1[1])) / float64(vec2[1])

	if math.Mod(b, 1) != 0 || math.Mod(a, 1) != 0 {
		return -1, -1
	}
	return int64(a), int64(b)
}

func main() {
	lines := utils.LoadFile(os.Args[1])
	machineInput := [4][2]int64{}
	firstResult := int64(0)
	secondResult := int64(0)
	for i, line := range lines {
		if i%4 == 0 || i%4 == 1 {
			machineInput[i%4] = splitLine(line, false)
		} else if i%4 == 2 {
			machineInput[i%4] = splitLine(line, false)
			machineInput[i%4+1] = splitLine(line, true)
		} else {
			a, b := FindCombination(machineInput[0], machineInput[1], machineInput[2])
			if a != -1 && b != -1 {
				firstResult += 3*a + b
			}
			a, b = FindCombination(machineInput[0], machineInput[1], machineInput[3])
			if a != -1 && b != -1 {
				secondResult += 3*a + b
			}
		}
	}
	a, b := FindCombination(machineInput[0], machineInput[1], machineInput[2])
	if a != -1 && b != -1 {
		firstResult += 3*a + b
	}
	a, b = FindCombination(machineInput[0], machineInput[1], machineInput[3])
	if a != -1 && b != -1 {
		secondResult += 3*a + b
	}
	fmt.Println("Result for the first part is: ", firstResult)
	fmt.Println("Result for the second part is: ", secondResult)
}
