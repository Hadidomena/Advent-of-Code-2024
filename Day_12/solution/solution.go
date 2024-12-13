package main

import (
	"fmt"
	"os"
	"strings"
)

type Queue [][]int

func (q *Queue) Enqueue(value []int) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() ([]int, bool) {
	if len(*q) == 0 {
		return []int{-1, -1}, false // Queue is empty
	}
	value := (*q)[0]
	*q = (*q)[1:] // Remove the front element
	return value, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
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
			if input[newX][newY] == value { // Check if the value equals previous one
				result = append(result, []int{newX, newY})
			}
		}
	}

	return result
}
func allBordering(input [][]rune, position [2]int) [][]int {
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
			if input[newX][newY] != value { // Check if the value equals previous one
				result = append(result, []int{newX, newY})
			}
		} else {
			result = append(result, []int{newX, newY})
		}
	}

	return result
}
func solution(board [][]rune) []int {
	allVisited := make(map[[2]int]bool)
	firstResult := 0
	secondResult := 0
	var q Queue
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			visited := make(map[[2]int]bool)
			start := [2]int{x, y}
			if _, exists := allVisited[start]; exists {
				continue // Skip if already processed
			}
			q.Enqueue([]int{x, y})
			visited[start] = true
			allVisited[start] = true
			perimeter := 0
			area := 1
			for !q.IsEmpty() {
				perimeter += 4
				node, _ := q.Dequeue()
				bordering := getBordering(board, node)
				for _, point := range bordering {
					perimeter--
					pointKey := [2]int{point[0], point[1]}
					if !allVisited[pointKey] {
						allVisited[pointKey] = true
						visited[pointKey] = true
						q.Enqueue(point)
						area += 1
					}
				}
			}

			borders := [][]int{}
			for point := range visited {
				borders = append(borders, allBordering(board, point)...)
			}

			directions := [][]int{
				{-1, 0}, // Up
				{1, 0},  // Down
				{0, -1}, // Left
				{0, 1},  // Right
			}
			allEdgePoints := make(map[[3]int]struct{})
			for _, point := range borders {
				for _, dir := range directions {
					pointKey := [2]int{point[0] + dir[0], point[1] + dir[1]}
					if visited[pointKey] {
						if dir[0] == 0 {
							if dir[1] == -1 {
								allEdgePoints[[3]int{point[0], point[1], 0}] = struct{}{}
							} else {
								allEdgePoints[[3]int{point[0], point[1], 1}] = struct{}{}
							}
						}
						if dir[1] == 0 {
							if dir[0] == -1 {
								allEdgePoints[[3]int{point[0], point[1], 2}] = struct{}{}
							} else {
								allEdgePoints[[3]int{point[0], point[1], 3}] = struct{}{}
							}
						}
					}
				}
			}
			edges := 0
			for len(allEdgePoints) > 0 {
				var firstKey [3]int
				for k := range allEdgePoints {
					firstKey = k
					break // Stop after the first iteration
				}
				visitedPts := bfs(firstKey[:], allEdgePoints)
				for point := range visitedPts {
					delete(allEdgePoints, point)
				}
				edges++
			}
			firstResult += perimeter * area
			secondResult += area * edges
		}
	}
	return []int{firstResult, secondResult}
}
func neighborCreator(node []int) [][]int {
	directions := [][]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
	neighboors := [][]int{}
	for _, dir := range directions {
		neighboors = append(neighboors, []int{node[0] + dir[0], node[1] + dir[1], node[2]})
	}
	return neighboors
}
func neighborCheck(neighbor [3]int, allEdgePoints map[[3]int]struct{}) bool {
	_, found := allEdgePoints[neighbor]
	return found
}

func bfs(start []int, allEdgePoints map[[3]int]struct{}) map[[3]int]struct{} {
	visited := make(map[[3]int]struct{})
	var q Queue
	visited[[3]int{start[0], start[1], start[2]}] = struct{}{}
	q.Enqueue(start)

	for !q.IsEmpty() {
		cp, _ := q.Dequeue()
		for _, neighbor := range neighborCreator(cp) {
			var array [3]int
			copy(array[:], neighbor)
			if neighborCheck(array, allEdgePoints) {
				if _, found := visited[array]; !found {
					q.Enqueue(neighbor)
					visited[array] = struct{}{}
				}
			}
		}
	}

	return visited
}

func main() {
	board := convertToRuneBoard(loadFile(os.Args[1]))

	result := solution(board)
	fmt.Println("Result for part 1 is ", result[0])
	fmt.Println("Result for part 2 is ", result[1])
}
