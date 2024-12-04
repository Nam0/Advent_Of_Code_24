package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D4input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read input into a grid
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Search word
	word := "XMAS"
	// coming back to this after pt 2 I feel like I should note I thought we'd have to do it with other words
	// which is why it's not a static coded word could've saved my self a bit of logic alwell

	// find em
	results := findWord(grid, word)
	total := 0

	// go thru results
	if len(results) == 0 {
		fmt.Println("fucky wucky")
	} else {
		fmt.Println("Words!")
		for _, result := range results {
			fmt.Printf("Start: (%d, %d), Direction: %s\n", result.row, result.col, result.direction)
			total++
		}
	}
	fmt.Printf("Found it: %d\n", total)
}

type result struct {
	row, col  int
	direction string
}

func findWord(grid []string, word string) []result {
	var results []result
	rows := len(grid)
	if rows == 0 {
		return results
	}
	cols := len(grid[0])

	// row delta col delta direction name
	directions := []struct {
		dRow, dCol int
		name       string
	}{
		{0, 1, "R"},
		{0, -1, "L"},
		{1, 0, "D"},
		{-1, 0, "U"},
		{1, 1, "Diag DR"},
		{-1, -1, "Diag UL"},
		{1, -1, "Diag DL"},
		{-1, 1, "Diag UP"},
	}

	// go thru every cell
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if checkWord(grid, word, r, c, dir.dRow, dir.dCol) {
					results = append(results, result{row: r, col: c, direction: dir.name})
				}
			}
		}
	}

	return results
}

func checkWord(grid []string, word string, startRow, startCol, dRow, dCol int) bool {
	for i := 0; i < len(word); i++ {
		r := startRow + i*dRow
		c := startCol + i*dCol

		// check bounds
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return false
		}

		// check char match
		if grid[r][c] != word[i] {
			return false
		}
	}
	return true
}
