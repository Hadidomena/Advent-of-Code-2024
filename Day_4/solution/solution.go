package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func secondPart(input []string) int {
	lenLine := len(input[0]) - 1
	amountOfLines := len(input)
	result := 0
	for i := 1; i < amountOfLines-1; i++ {
		for j := 1; j < lenLine-1; j++ {
			if input[i][j] == 'A' {
				if ((input[i+1][j+1] == 'M' && input[i-1][j-1] == 'S') ||
					(input[i+1][j+1] == 'S' && input[i-1][j-1] == 'M')) &&
					((input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S') ||
						(input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M')) {
					result += 1
				}
			}
		}
	}
	return result
}
func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Błąd wczytywania pliku:", err)
		return
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	lenLine := len(lines[0]) - 1
	amountOfLines := len(lines) - 1

	var creator strings.Builder
	for i := 0; i < lenLine; i++ {
		for j := 0; j < amountOfLines+1; j++ {
			creator.WriteByte(lines[j][i])
		}
		creator.WriteByte('\n')
	}
	columns := creator.String()

	creator.Reset()
	for i := 0; i < amountOfLines; i++ {
		row, col := i, 0 // Start from row i, first column
		for row >= 0 && col < lenLine {
			creator.WriteByte(lines[row][col])
			row-- // Move upwards
			col++ // Move to the next column
		}
		creator.WriteByte('\n') // Add a newline after each diagonal
	}
	for i := 0; i < lenLine; i++ {
		row, col := amountOfLines, i // Start from the last row, column i
		for row >= 0 && col < lenLine {
			creator.WriteByte(lines[row][col])
			row-- // Move upwards
			col++ // Move to the next column
		}
		creator.WriteByte('\n') // Add a newline after each diagonal
	}
	diagonalLeft := creator.String()

	creator.Reset()
	for i := 0; i < amountOfLines; i++ {
		row, col := i, lenLine-1 // Start from row i, last column
		for row >= 0 && col >= 0 {
			creator.WriteByte(lines[row][col])
			row-- // Move upwards
			col-- // Move to the previous column
		}
		creator.WriteByte('\n') // Add a newline after each diagonal
	}

	// Second pass: Start from the first column of the last row
	for i := lenLine - 1; i >= 0; i-- { // Start from second last column to avoid duplicate diagonal
		row, col := amountOfLines, i // Start from the last row, column i
		for row > 0 && col >= 0 {
			creator.WriteByte(lines[row][col])
			row-- // Move upwards
			col-- // Move to the previous column
		}
		creator.WriteByte('\n') // Add a newline after each diagonal
	}
	diagonalsRight := creator.String()

	result := 0
	result += check(diagonalLeft)
	result += check(diagonalsRight)
	result += check(content)
	result += check(columns)

	fmt.Println(result)
	fmt.Println(secondPart(lines))

}

func check(input string) int {
	result := 0
	pattern := `XMAS`
	patternBackwards := `SAMX`

	re := regexp.MustCompile(pattern)
	reb := regexp.MustCompile(patternBackwards)

	matches := re.FindAllString(input, -1)
	result += len(matches)
	matches = reb.FindAllString(input, -1)
	result += len(matches)

	return result
}

/*creator.Reset()
// Left diagonals with unique path generation
for startRow := 0; startRow <= amountOfLines; startRow++ {
	var diagonal []byte
	row, col := startRow, 0
	for k := 0; k < lenLine; k++ {
		diagonal = append(diagonal, lines[row][col])
		row = (row - 1 + amountOfLines + 1) % (amountOfLines + 1)
		col = (col + 1) % lenLine
	}
	creator.Write(diagonal)
	creator.WriteByte('\n')
}
diagonalLeft := creator.String()

creator.Reset()
// Right diagonals with unique path generation
for startRow := 0; startRow <= amountOfLines; startRow++ {
	var diagonal []byte
	row, col := startRow, lenLine-1
	for k := 0; k < lenLine; k++ {
		fmt.Println(row, col)
		diagonal = append(diagonal, lines[row][col])
		row = (row - 1 + amountOfLines + 1) % (amountOfLines + 1)
		col = (col - 1 + lenLine) % lenLine
	}
	creator.Write(diagonal)
	creator.WriteByte('\n')
}
diagonalsRight := creator.String()*/
