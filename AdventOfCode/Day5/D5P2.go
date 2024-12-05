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
		} else if strings.Contains(line, ",") { // Update lines
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

// func to put it in the correct order based on rule set
func CorrectOrder(update []int, rules []Rule) []int {
	// copy og slice so we dont have to modify it
	pages := append([]int(nil), update...)

	// change until we dont need to
	for changed := true; changed; {
		changed = false // reset flag
		// check for invalid on what rule
		for _, rule := range rules {
			beforePage, afterPage := -1, -1
			// Find pos of the before and after pages
			for k, page := range pages {
				if page == rule.before {
					beforePage = k
				} else if page == rule.after {
					afterPage = k
				}
			}
			// if both pages found and before comes after the after page swap their pos
			if beforePage > afterPage && beforePage != -1 && afterPage != -1 {
				pages[beforePage], pages[afterPage] = pages[afterPage], pages[beforePage]
				changed = true // mark it as changed
			}
		}
	}

	return pages
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
		if !IsValidOrder(update, rules) {
			// Correct the order of the update
			correctedUpdate := CorrectOrder(update, rules)
			// Add the middle page to the total
			sumOfMiddlePages += correctedUpdate[len(correctedUpdate)/2]
		}
	}

	fmt.Printf("Sum: %d\n", sumOfMiddlePages)
}
