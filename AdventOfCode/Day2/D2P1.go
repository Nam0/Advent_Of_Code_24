package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(levels []int) bool {
	// Check if dif is between 1 and 3
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
