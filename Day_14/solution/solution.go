package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Hadidomena/AoC_Go_utils/utils"
)

type robot struct {
	position [2]int
	speed    [2]int
}

func (r *robot) pushRobot(howMuch, boardWidth, boardHeight int) {
	r.position[0] = (r.position[0] + howMuch*r.speed[0]) % boardWidth
	r.position[1] = (r.position[1] + howMuch*r.speed[1]) % boardHeight

	if r.position[0] < 0 {
		r.position[0] += boardWidth
	}
	if r.position[1] < 0 {
		r.position[1] += boardHeight
	}
}
func splitLine(input string) robot {
	split := strings.Split(input, " ")
	startingPosition, speed := [2]int{}, [2]int{}
	for i, element := range split {
		element = element[2:]
		numbers := strings.Split(element, ",")
		for j, number := range numbers {
			if i == 0 {
				num, err := strconv.Atoi(number)
				if err == nil {
					startingPosition[j] = num
				}
			} else {
				num, err := strconv.Atoi(number)
				if err == nil {
					speed[j] = num
				}
			}
		}
	}
	return robot{position: startingPosition, speed: speed}
}
func tryTree(x, boardWidth, boardHeight int) {
	lines := utils.LoadFile(os.Args[1])
	var q utils.Queue[robot]
	for _, line := range lines {
		robot := splitLine(line)
		robot.pushRobot(x, boardWidth, boardHeight)
		q.Enqueue(robot)
	}

	frequencies := make(map[[2]int]int)
	for !q.IsEmpty() {
		robot, _ := q.Dequeue()
		frequencies[robot.position]++
	}

	tmpFileName := "board_state.txt"
	file, err := os.OpenFile(tmpFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write board state to the file
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pos := [2]int{x, y}
			if _, exists := frequencies[pos]; exists {
				fmt.Fprint(file, "#") // Write '#' for a point with frequency
			} else {
				fmt.Fprint(file, ".") // Write '.' for empty
			}
		}
		fmt.Fprintln(file) // Newline for next row
	}

	// Write the x value at the end
	fmt.Fprintln(file, x)
}

func partTwo(boardWidth, boardHeight int) {
	x := 0
	for x < 204967 {
		tryTree(x, boardWidth, boardHeight)
		x++
	}

}

func partOne(boardWidth, boardHeight int) {
	lines := utils.LoadFile(os.Args[1])
	var q utils.Queue[robot]
	for _, line := range lines {
		robot := splitLine(line)
		robot.pushRobot(100, boardWidth, boardHeight)
		q.Enqueue(robot)
	}

	quarters := [4]int{}
	for !q.IsEmpty() {
		robot, _ := q.Dequeue()
		if robot.position[0] < boardWidth/2 && robot.position[1] < boardHeight/2 {
			quarters[0]++
		} else if robot.position[0] > boardWidth/2 && robot.position[1] < boardHeight/2 {
			quarters[1]++
		} else if robot.position[0] > boardWidth/2 && robot.position[1] > boardHeight/2 {
			quarters[2]++
		} else if robot.position[0] < boardWidth/2 && robot.position[1] > boardHeight/2 {
			quarters[3]++
		}
	}

	fmt.Println("Result for the first part is: ", quarters[0]*quarters[1]*quarters[2]*quarters[3])
}
func main() {
	partOne(101, 103)
	partTwo(101, 103)
}
