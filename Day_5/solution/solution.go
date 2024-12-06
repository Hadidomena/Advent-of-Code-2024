package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	befores := make(map[string][]string)
	isRecords := false
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Błąd wczytywania pliku:", err)
		return
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	result := 0
	secondResult := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			isRecords = true
		} else {
			if !isRecords {
				// Handle dependencies
				splitLine := strings.Split(strings.TrimSpace(line), "|")
				addToMap(strings.TrimSpace(splitLine[0]), strings.TrimSpace(splitLine[1]), befores)
			} else {
				// Handle sequences
				splitLine := strings.Split(line, ",") // Split on commas for sequences
				for i,j := range splitLine {
					splitLine[i] = strings.TrimSpace(j)
				}
				if validateOrder(splitLine, befores) {
					//fmt.Println(line) // Only print valid sequences
					num, _ := strconv.Atoi(splitLine[len(splitLine) / 2])
					result += num
					//fmt.Println(num)
				} else {
					splitLine = correctSequence(splitLine, befores)
					num, _ := strconv.Atoi(splitLine[len(splitLine) / 2])
					secondResult += num
				}
			}
		}
	}
	fmt.Println(result)
	fmt.Println(secondResult)
}

func addToMap(key string, value string, data map[string][]string) {
	if _, exists := data[key]; exists {
		data[key] = append(data[key], value)
	} else {
		data[key] = []string{value}
	}
}

func validateOrder(sequence []string, befores map[string][]string) bool {
    // Create a graph to track dependencies
    graph := make(map[string]map[string]bool)
    
    // Build the dependency graph
    for page, deps := range befores {
        if graph[page] == nil {
            graph[page] = make(map[string]bool)
        }
        for _, dep := range deps {
            graph[page][dep] = true
        }
    }
    
    // Check the order of the sequence
    for i := 0; i < len(sequence); i++ {
        for j := i + 1; j < len(sequence); j++ {
            // Check if there's a dependency rule that would be violated
            if graph[sequence[j]] != nil && 
               graph[sequence[j]][sequence[i]] == true {
                return false
            }
        }
    }
    
    return true
}

func correctSequence(sequence []string, befores map[string][]string) []string {
    // Create a graph to track dependencies
    graph := make(map[string]map[string]bool)
    
    // Build the dependency graph
    for page, deps := range befores {
        if graph[page] == nil {
            graph[page] = make(map[string]bool)
        }
        for _, dep := range deps {
            graph[page][dep] = true
        }
    }
    
    // Create a copy of the original sequence to modify
    corrected := make([]string, len(sequence))
    copy(corrected, sequence)
    
    // Bubble sort with dependency-aware swapping
    for i := 0; i < len(corrected); i++ {
        for j := 0; j < len(corrected)-i-1; j++ {
            // Check if a swap is needed due to dependencies
            if graph[corrected[j+1]] != nil && 
               graph[corrected[j+1]][corrected[j]] == true {
                // Swap the elements
                corrected[j], corrected[j+1] = corrected[j+1], corrected[j]
            }
        }
    }
    
    // Optional: Add a final validation check
    if validateOrder(corrected, befores) {
        return corrected
    }
    
    // If no valid correction found, return original sequence
    return sequence
}