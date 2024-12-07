package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func to evaluate left to right w either + * or || to combind
func evaluate(numbers []int, operators []string) int {
	result := numbers[0]
	for i, op := range operators {
		if op == "+" {
			result += numbers[i+1]
		} else if op == "*" {
			result *= numbers[i+1]
		} else if op == "||" {
			// concat as string then change to int
			concatStr := fmt.Sprintf("%d%d", result, numbers[i+1])
			concatValue, _ := strconv.Atoi(concatStr)
			result = concatValue
		}
	}
	return result
}

// func to generate all possible combos
func generateOperatorCombinations(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	if n == 1 {
		return [][]string{{"+"}, {"*"}, {"||"}}
	}

	subCombinations := generateOperatorCombinations(n - 1)
	var combinations [][]string

	for _, sub := range subCombinations {
		combinations = append(combinations, append([]string{"+"}, sub...))
		combinations = append(combinations, append([]string{"*"}, sub...))
		combinations = append(combinations, append([]string{"||"}, sub...))
	}
	return combinations
}

func main() {
	fileName := "D7input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		// read test value
		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Printf("Invalid test value in line: %s\n", line)
			continue
		}

		// read numbers
		numStrs := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			numbers[i], err = strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Invalid number in line: %s\n", line)
				continue
			}
		}

		// gen all combos
		operatorCombinations := generateOperatorCombinations(len(numbers) - 1)

		// check combos
		valid := false
		for _, operators := range operatorCombinations {
			if evaluate(numbers, operators) == testValue {
				valid = true
				break
			}
		}

		// add to total if valid
		if valid {
			total += testValue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Total : %d\n", total)
}
