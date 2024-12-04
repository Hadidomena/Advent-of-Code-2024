package main

import (
	"os"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)
func solution() {
	// Wczytaj plik
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Błąd wczytywania pliku:", err)
		return
	}
	content := string(data)
	// Regular expression to match mul(x,y) pattern
	pattern := `mul\(\d{1,3},\d{1,3}\)`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Split the text into sections based on "don't()"
	sections := strings.Split(content, "don't()")

	// Container for results
	results := []string{}

	// Process the first section (from the beginning of the string to the first "don't()")
	if len(sections) > 0 {
		firstSection := sections[0]
		matches := re.FindAllString(firstSection, -1)
		results = append(results, matches...)
	}

	// Process subsequent sections for "do() ... don't()" ranges
	for i := 1; i < len(sections); i++ {
		// For each section, split by "do()" and process relevant parts
		subSections := strings.Split(sections[i], "do()")
		for j := 1; j < len(subSections); j++ {
			// Match within the part after "do()" and before "don't()"
			matches := re.FindAllString(subSections[j], -1)
			results = append(results, matches...)
		}
	}

	result := 0
	for _, match := range results {
		result += mul(match)
	}
	fmt.Println("Wynik =", result)
}
func mul(input string) int {
	subset := input[4:len(input) - 1]
	elements := strings.Split(subset, ",")
	intElements := []int{}

	for _, element := range elements {
		// Konwertuj każdy element na int
		num, err := strconv.Atoi(strings.TrimSpace(element))
		if err != nil {
			fmt.Println("Błąd konwersji:", err)
			continue
		}
		intElements = append(intElements, num)
	}

	return intElements[0] * intElements[1]
}

func main() {
	solution()
}