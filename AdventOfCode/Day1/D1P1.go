package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// Sort lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Find totalDistance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
}
