package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileName := "sanitized_input"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// regex  my beloved
	// yes I cheated I love https://www.rexegg.com/regex-quickstart.php so much and cba to remember it
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRegex := regexp.MustCompile(`\bdo\(\)`) // these ones are easier tho
	dontRegex := regexp.MustCompile(`\bdon't\(\)`)

	total := 0
	mulEnabled := true // start enabled

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("pLine: %s mulEnabled=%v\n", line, mulEnabled)

		// Handle do() and don't()
		if doRegex.MatchString(line) {
			fmt.Println("do()")
			mulEnabled = true
		}
		if dontRegex.MatchString(line) {
			fmt.Println("dont()")
			mulEnabled = false
		}

		//Attempted to use this originally however due to bad implementation it didn't work, because the data was v dirty
		// and idk what Im doing, wrote D3P3V2 to sanitize the data only keep dos donts and muls(regardless of numbers) and then process it here.
		if mulEnabled {
			matches := mulRegex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				// conv to int
				x, err1 := strconv.Atoi(match[1])
				y, err2 := strconv.Atoi(match[2])
				if err1 == nil && err2 == nil {
					fmt.Printf("Multiplying %d * %d\n", x, y)
					total += x * y
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Total: %d\n", total)
}
