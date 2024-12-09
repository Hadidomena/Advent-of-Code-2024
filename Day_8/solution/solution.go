package main

import (
	"fmt"
	"os"
	"strings"
)

var maxX int
var maxY int

func outsideRange(input []int) bool {
	return input[0] > maxY || input[0] < 0 || input[1] > maxX || input[1] < 0
}
func solution1(input map[rune][][]int) int {
	result := 0
	uniqueLocations := make(map[string]bool) // Track unique locations

	for _, positions := range input {
		for i, position1 := range positions {
			for j, position2 := range positions {
				if j > i {
					diffX := position2[0] - position1[0]
					diffY := position2[1] - position1[1]

					// First created position
					created := []int{position1[0] - diffX, position1[1] - diffY}
					createdKey := fmt.Sprintf("%d,%d", created[0], created[1]) // Convert to string key
					if !outsideRange(created) && !uniqueLocations[createdKey] {
						uniqueLocations[createdKey] = true
						result++
					}

					// Second created position
					created2 := []int{position2[0] + diffX, position2[1] + diffY}
					created2Key := fmt.Sprintf("%d,%d", created2[0], created2[1]) // Convert to string key
					if !outsideRange(created2) && !uniqueLocations[created2Key] {
						uniqueLocations[created2Key] = true
						result++
					}
				}
			}
		}
	}

	return result
}
func solution2(input map[rune][][]int) int {
	result := 0
	uniqueLocations := make(map[string]bool) // Track unique locations

	for _, positions := range input {
		for i, position1 := range positions {
			for j, position2 := range positions {
				if j > i {
					diffX := position2[0] - position1[0]
					diffY := position2[1] - position1[1]

					iteration := 0
					for {
						// First created row
						created := []int{position1[0] - diffX*iteration, position1[1] - diffY*iteration}
						createdKey := fmt.Sprintf("%d,%d", created[0], created[1]) // Convert to string key
						if outsideRange(created) {
							break
						} else if !uniqueLocations[createdKey] {
							uniqueLocations[createdKey] = true
							result++
						}
						iteration++
					}

					iteration = 0
					for {
						created2 := []int{position2[0] + diffX*iteration, position2[1] + diffY*iteration}
						created2Key := fmt.Sprintf("%d,%d", created2[0], created2[1]) // Convert to string key
						if outsideRange(created2) {
							break
						} else if !uniqueLocations[created2Key] {
							uniqueLocations[created2Key] = true
							result++
						}
						iteration++
					}
				}
			}
		}
	}

	return result
}
func main() {
	var positions = make(map[rune][][]int)
	// Read the input file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := string(data)
	lines := strings.Split(content, "\r\n")

	maxY = len(lines) - 1
	maxX = len(lines[0]) - 1

	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				if _, exists := positions[char]; exists {
					positions[char] = append(positions[char], []int{i, j})
				} else {
					positions[char] = [][]int{{i, j}}
				}
			}
		}
	}

	echoPoints := solution1(positions)
	fmt.Println("Result 1: ", (echoPoints))

	echoPoints = solution2(positions)
	fmt.Println("Result 2: ", (echoPoints))
}
