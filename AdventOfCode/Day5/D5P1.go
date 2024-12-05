package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// X|Y X then Y
type Rule struct {
	before int
	after  int
}

// func to parse thru input file for rule lines and update lines
func ParseRulesAndUpdates(fileName string) ([]Rule, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var rules []Rule
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") { // Rule lines
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule{before: before, after: after})
		} else if strings.Contains(line, ",") { //Update lines
			pageStrings := strings.Split(line, ",")
			var pages []int
			for _, ps := range pageStrings {
				page, _ := strconv.Atoi(ps)
				pages = append(pages, page)
			}
			updates = append(updates, pages)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return rules, updates, nil
}

// func to check if the update line is valid
func IsValidOrder(update []int, rules []Rule) bool {
	// make map for lookup
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	// itterate thru and check
	for _, rule := range rules {
		if posBefore, okBefore := position[rule.before]; okBefore {
			if posAfter, okAfter := position[rule.after]; okAfter {
				if posBefore >= posAfter {
					return false // doesn't follow rules
				}
			}
		}
	}

	return true
}

func main() {
	fileName := "D5input.txt"

	// Parse rules and updates
	rules, updates, err := ParseRulesAndUpdates(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	sumOfMiddlePages := 0
	for _, update := range updates {
		if IsValidOrder(update, rules) {
			middlePage := update[len(update)/2]
			sumOfMiddlePages += middlePage
		}
	}

	fmt.Printf("Sum: %d\n", sumOfMiddlePages)
}
