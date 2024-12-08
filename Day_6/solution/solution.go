package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var directions = map[string]int{
	"^": 0,
	">": 1,
	"v": 2,
	"<": 3,
}
var arr1 = []int{1, 2, 3}
var guard = &Guard{
	direction: 0,
	position:  arr1,
}
var initialPosition = []int{49, 47}
var initialDirection int = 0

// Guard struct to represent a guard with a direction and position
type Guard struct {
	direction int
	position  []int
}

// Function to check if a key is in the map
func isIn(element string, dirMap map[string]int) bool {
	_, exists := dirMap[element]
	return exists
}

// Function to convert a string board to a mutable rune board
func convertToRuneBoard(board []string) [][]rune {
	runeBoard := make([][]rune, len(board))
	for i, line := range board {
		runeBoard[i] = []rune(line)
	}
	return runeBoard
}

// Function to convert a rune board back to a string board
func convertToStringBoard(runeBoard [][]rune) []string {
	stringBoard := make([]string, len(runeBoard))
	for i, line := range runeBoard {
		stringBoard[i] = string(line)
	}
	return stringBoard
}

func moveAhead(board [][]rune) int {
	position := guard.position
	if guard.direction == 0 {
		for {
			if board[position[0]][position[1]] == '#' {
				position[0]++
				guard.position = position
				guard.direction = (guard.direction + 1) % 4 // Turn up
				return 1
			} else if position[0] == 0 {
				board[position[0]][position[1]] = 'X'
				return 0
			} else {
				board[position[0]][position[1]] = 'X'
			}
			position[0]--
		}
	} else if guard.direction == 1 {
		for {
			if board[position[0]][position[1]] == '#' {
				position[1]--
				guard.position = position
				guard.direction = (guard.direction + 1) % 4 // Turn up
				return 1
			} else if position[1] == len(board[0])-1 {
				board[position[0]][position[1]] = 'X'
				return 0
			} else {
				board[position[0]][position[1]] = 'X'
			}
			position[1]++
		}
	} else if guard.direction == 2 {
		for {
			if board[position[0]][position[1]] == '#' {
				position[0]--
				guard.position = position
				guard.direction = (guard.direction + 1) % 4 // Turn up
				return 1
			} else if position[0] == len(board)-1 {
				board[position[0]][position[1]] = 'X'
				return 0
			} else {
				board[position[0]][position[1]] = 'X'
			}
			position[0]++
		}
	} else if guard.direction == 3 {
		for {
			if board[position[0]][position[1]] == '#' {
				position[1]++
				guard.position = position
				guard.direction = (guard.direction + 1) % 4 // Turn up
				return 1
			} else if position[1] == 0 {
				board[position[0]][position[1]] = 'X'
				return 0
			} else {
				board[position[0]][position[1]] = 'X'
			}
			position[1]--
		}
	}
	return 2
}

func main() {
	// Read the input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the file content
	content := string(data)
	lines := strings.Split(content, "\n")

	// Convert lines to a mutable rune board
	board := convertToRuneBoard(lines)
	copiedBoard := make([][]rune, len(board))
	for i := range board {
		copiedBoard[i] = make([]rune, len(board[i]))
		copy(copiedBoard[i], board[i])
	}

	// Find the guard's position and direction
	for x := 0; x < len(board); x++ {
		for y, value := range board[x] {
			char := string(value) // Convert rune to string
			if isIn(char, directions) {
				position := []int{x, y}
				guard.position = position
				guard.direction = directions[char]
				break
			}
		}
	}

	// Print the guard's details
	if guard != nil {
		fmt.Printf("Guard found at position %v with direction %d\n", guard.position, guard.direction)
	} else {
		fmt.Println("No guard found in the input.")
		return
	}

	// Move the guard until it stops
	for {
		result := moveAhead(board)
		if result == 0 {
			break
		}
	}
	// Regular expression to match mul(x,y) pattern
	pattern := `X`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Split the text into sections based on "don't()"
	// Convert the rune board back to a string board and print the final state
	traveledTiles := 0
	finalBoard := convertToStringBoard(board)
	for _, line := range finalBoard {
		content := string(line)
		matches := re.FindAllString(content, -1)
		traveledTiles += len(matches)
	}
	fmt.Println("It traveled through: ", traveledTiles)
	fmt.Println("You can place obstacles: ", len(findTrapPositions(copiedBoard)))
}

func findTrapPositions(input [][]rune) [][]int {
	trapPositions := [][]int{}
	// Iterate through all positions on the board
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[x]); y++ {
			board := make([][]rune, len(input))
			for i := range input {
				board[i] = make([]rune, len(input[i]))
				copy(board[i], input[i])
			}

			if board[x][y] == '.' { // Consider empty spaces only
				// Temporarily place an obstacle
				board[x][y] = '#'

				// Check if this causes a cycle
				if isCycle(board) {
					trapPositions = append(trapPositions, []int{x, y})
				}

				// Remove the obstacle
				board[x][y] = '.'
			}
		}
	}

	return trapPositions
}

// Helper function to determine if the guard becomes trapped in a cycle
func isCycle(board [][]rune) bool {
	// Create a map to track visited states: (position, direction)
	visited := make(map[string]bool)
	// Start with the initial guard position and direction
	copy(guard.position, initialPosition)
	guard.direction = initialDirection

	for {
		// Encode the current state as a unique string
		state := fmt.Sprintf("%d,%d,%d", guard.position[0], guard.position[1], guard.direction)
		state = strings.TrimSpace(state)
		// If this state has already been visited, a cycle is detected
		_, exists := visited[state]
		if exists {
			return exists
		} else {
			// Mark the current state as visited
			visited[state] = true
		}

		// Move the guard ahead
		result := moveAhead(board)
		// If the guard stops, there's no cycle
		if result == 0 {
			return false
		}
	}
}
