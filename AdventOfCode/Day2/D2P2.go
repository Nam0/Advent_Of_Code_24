package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(levels []int) bool {
	checkSafe := func(levels []int) bool {
		// Check if diff between 1 and 3
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]
			if diff < -3 || diff > 3 || diff == 0 {
				return false
			}
		}

		// Check if sonly increasing or dec
		isIncreasing := true
		isDecreasing := true
		for i := 1; i < len(levels); i++ {
			if levels[i] > levels[i-1] {
				isDecreasing = false
			}
			if levels[i] < levels[i-1] {
				isIncreasing = false
			}
		}

		return isIncreasing || isDecreasing
	}

	// Check if the report is safe right away
	if checkSafe(levels) {
		return true
	}

	// Ittr thru levels and see if it iwll return true with it removed
	for i := 0; i < len(levels); i++ {
		// Make a list with all elements but i
		modified := append([]int{}, levels[:i]...)
		modified = append(modified, levels[i+1:]...)
		if checkSafe(modified) {
			return true
		}
	}

	// More than one removal required
	return false
}

func main() {
	fileName := "D2input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into numbers
		numberStrings := strings.Fields(line) // I LOVE FIELDS RAHHHHHHHHH
		var levels []int

		for _, numStr := range numberStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Skipping invalid number: %s\n", numStr)
				continue
			}
			levels = append(levels, num)
		}

		// check if is safe
		if isSafeReport(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Total safe reports: %d\n", safeCount)
}

func test() {
	testCases := [][]int{
		{7, 6, 4, 2, 1}, // Safe without ignoring
		{1, 2, 7, 8, 9}, // Unsafe even with ignoring
		{9, 7, 6, 2, 1}, // Unsafe even with ignoring
		{1, 3, 2, 4, 5}, // Safe by ignoring 2
		{8, 6, 4, 4, 1}, // Safe by ignoring 4
		{1, 3, 6, 7, 9}, // Safe without ignoring
	}

	for _, levels := range testCases {
		fmt.Printf("Levels: %v, Safe: %v\n", levels, isSafeReport(levels))
	}
}
