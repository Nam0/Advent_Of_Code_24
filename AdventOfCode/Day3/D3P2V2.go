package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputFile := "D3input.txt"
	outputFile := "sanitized_input"

	// simplified expressions
	validPatterns := `do\(\)|don't\(\)|mul\(\d+,\d+\)`

	// Compile the regex
	regex, err := regexp.Compile(validPatterns)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return
	}

	// in
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inFile.Close()

	// out
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outFile.Close()

	// go by line by line for all 6 lines
	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		// matches any of the regex
		matches := regex.FindAllString(line, -1)
		if matches != nil {
			for _, match := range matches {
				_, err := writer.WriteString(match + "\n")
				if err != nil {
					fmt.Printf("Error writing to output file: %v\n", err)
					return
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
}
