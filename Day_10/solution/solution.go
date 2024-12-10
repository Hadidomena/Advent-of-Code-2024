package main

import (
	"fmt"
	"os"
	"strings"
)

type Queue [][]int

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value []int) {
	*q = append(*q, value)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() ([]int, bool) {
	if len(*q) == 0 {
		return []int{-1, -1}, false // Queue is empty
	}
	value := (*q)[0]
	*q = (*q)[1:] // Remove the front element
	return value, true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func convertToRuneBoard(board []string) [][]rune {
	runeBoard := make([][]rune, len(board))
	for i, line := range board {
		runeBoard[i] = []rune(line)
	}
	return runeBoard
}

func getBordering(input [][]rune, position []int) [][]int {
	value := input[position[0]][position[1]]
	result := [][]int{}

	// Define the directions for neighboring positions (up, down, left, right)
	directions := [][]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}

	// Iterate through each direction
	for _, direction := range directions {
		newX := position[0] + direction[0]
		newY := position[1] + direction[1]

		// Check if the new position is within bounds
		if newX >= 0 && newX < len(input) && newY >= 0 && newY < len(input[0]) {
			if input[newX][newY] == value+1 { // Check if the value equals `value + 1`
				result = append(result, []int{newX, newY})
			}
		}
	}

	return result
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := string(data)
	lines := strings.Split(content, "\r\n")
	board := convertToRuneBoard(lines)

	// Map to track visited points for each starting `0`
	visited := make(map[[2]int]map[[2]int]bool)

	result := 0

	// Initialize a new queue and visited map for this `0`
	var q Queue

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == '0' {
				start := [2]int{x, y}
				if _, exists := visited[start]; exists {
					continue // Skip if already processed
				}
				q.Enqueue([]int{x, y})
				visited[start] = make(map[[2]int]bool)
				visited[start][start] = true

				// Process BFS for the current `0`
				for !q.IsEmpty() {
					node, _ := q.Dequeue()
					if board[node[0]][node[1]] == '9' {
						result++
					} else {
						bordering := getBordering(board, node)
						for _, point := range bordering {
							pointKey := [2]int{point[0], point[1]}
							if !visited[start][pointKey] {
								visited[start][pointKey] = true
								q.Enqueue(point)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("Result for part 1 is ", result)

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == '0' {
				q.Enqueue([]int{x, y})
			}
		}
	}
	fmt.Println("Result for part 2 is ", secondPart(q, board))
}
func secondPart(q Queue, board [][]rune) int {
	result := 0
	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		if board[node[0]][node[1]] == '9' {
			result++
		} else {
			bordering := getBordering(board, node)
			for _, point := range bordering {
				q.Enqueue(point)
			}
		}
	}
	return result
}
