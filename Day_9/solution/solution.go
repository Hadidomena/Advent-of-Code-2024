package main

import (
	"fmt"
	"os"
	"strconv"
)

func getToLeft(input []int64) {
	left := 0
	right := len(input) - 1

	for left < right {
		// Move `left` pointer to the next `-1`
		for left < right && input[left] != -1 {
			left++
		}

		// Move `right` pointer to the next non-`-1`
		for left < right && input[right] == -1 {
			right--
		}

		// Swap the elements if `left` is still less than `right`
		if left < right {
			input[left], input[right] = input[right], -1
		}
	}
}
func sumFromLeft(input []int64) int64 {
	result, _ := strconv.ParseInt("0", 10, 64)
	for index, value := range input {
		if value != -1 {
			result += int64(index) * value
		}
	}
	return result
}

func getBlocksToLeft(input []int64) {
	right := len(input) - 1

	for right >= 0 {
		// Move `right` pointer to the end of the next block (non-`-1`)
		for right >= 0 && input[right] == -1 {
			right--
		}

		// If no more blocks are found, we're done
		if right < 0 {
			break
		}

		// Find the contiguous block of the same value starting at `right`
		blockStart := right
		for blockStart > 0 && input[blockStart-1] == input[right] {
			blockStart--
		}
		blockSize := right - blockStart + 1

		// Attempt to find enough free spaces to fit the block starting from the beginning
		freeSpaceStart := 0
		for freeSpaceStart < blockStart {
			// Count the free spaces starting from `freeSpaceStart`
			freeSpaces := 0
			for i := freeSpaceStart; i < len(input) && input[i] == -1; i++ {
				freeSpaces++
			}

			// If the block fits, move it to the free space
			if freeSpaces >= blockSize {
				copy(input[freeSpaceStart:freeSpaceStart+blockSize], input[blockStart:right+1]) // Move block
				for i := blockStart; i <= right; i++ {                                          // Fill the old block with `-1`
					input[i] = -1
				}
				break
			}

			// Move `freeSpaceStart` to the end of the current free space segment
			for freeSpaceStart < len(input) && input[freeSpaceStart] == -1 {
				freeSpaceStart++
			}

			// Continue searching for the next free space
			for freeSpaceStart < len(input) && input[freeSpaceStart] != -1 {
				freeSpaceStart++
			}
		}

		// Move `right` pointer to process the next block
		right = blockStart - 1
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := string(data)
	numbers := []int64{}
	for index, sign := range content {
		if index%2 == 0 {
			max, _ := strconv.Atoi(string(sign))
			for i := 0; i < max; i++ {
				numbers = append(numbers, int64(index/2))
			}
		} else {
			max, _ := strconv.Atoi(string(sign))
			for i := 0; i < max; i++ {
				numbers = append(numbers, -1)
			}
		}
	}
	copiedNumbers := make([]int64, len(numbers))
	copy(copiedNumbers, numbers)

	getToLeft(numbers)
	fmt.Println("Result for part 1 is: ", sumFromLeft(numbers))

	getBlocksToLeft(copiedNumbers)
	fmt.Println("Result for part 2 is: ", sumFromLeft(copiedNumbers))

}
