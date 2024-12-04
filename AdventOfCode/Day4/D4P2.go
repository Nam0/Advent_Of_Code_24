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

	total := 0
	// find Mas
	masResults := findIntersectingMAS(grid)
	if len(masResults) == 0 {
		fmt.Println("fucky wucky")
	} else {
		fmt.Println("Yippe")
		for _, res := range masResults {
			fmt.Printf("Center A at (%d, %d)\n", res.row, res.col)
			total++
		}
	}
	fmt.Printf("Found it: %d\n", total)
}

type result struct {
	row, col  int
	direction string
}

func findIntersectingMAS(grid []string) []result {
	var results []result
	rows := len(grid)
	if rows == 0 {
		return results
	}
	cols := len(grid[0])

	// go thru each cell aside from edges
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] == 'A' && isIntersectingMAS(grid, r, c) {
				results = append(results, result{row: r, col: c})
			}
		}
	}

	return results
}

func isIntersectingMAS(grid []string, centerRow, centerCol int) bool {
	// dir struct
	directions := []struct {
		dRow, dCol int
	}{
		{-1, -1}, // UL diagonal
		{-1, 1},  // UR diagonal
		{1, -1},  // BL diagonal
		{1, 1},   // BR diagonal
	}

	found := 0
	for _, dir := range directions {
		// look for m and s diag from a
		mRow, mCol := centerRow+dir.dRow, centerCol+dir.dCol
		sRow, sCol := centerRow-dir.dRow, centerCol-dir.dCol

		if isInBounds(grid, mRow, mCol) && grid[mRow][mCol] == 'M' &&
			isInBounds(grid, sRow, sCol) && grid[sRow][sCol] == 'S' {
			found++
		}
	}

	// only counts if 2 are there
	return found >= 2
}

// func to make sure we dont try to reach OOB
func isInBounds(grid []string, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}
