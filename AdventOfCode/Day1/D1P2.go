package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "D1input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split line
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			fmt.Printf("Skipping invalid line: %s\n", line)
			continue
		}

		// Convert numbers to ints
		leftNum, err1 := strconv.Atoi(numbers[0])
		rightNum, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Skipping invalid numbers: %s\n", line)
			continue
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Count in right
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	// Find similarity score
	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	fmt.Printf("Similarity score: %d\n", similarityScore)
}
