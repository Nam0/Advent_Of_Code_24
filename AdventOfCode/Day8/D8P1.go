package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	fileName := "D8input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// map slices
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Add antennas into map by freq
	antennas := make(map[rune][]Point)
	for y, row := range grid {
		for x, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], Point{x, y})
			}
		}
	}

	// map to store antinode location
	antinodePositions := make(map[Point]struct{})

	// process each freq
	for _, points := range antennas {
		// comparry pairs of antennas with same freq
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				a, b := points[i], points[j]

				// calc vector between antennas
				dx, dy := b.x-a.x, b.y-a.y

				// set antinodes pos
				antinode1 := Point{a.x + 2*dx, a.y + 2*dy}
				antinode2 := Point{b.x - 2*dx, b.y - 2*dy}

				// if new antinode is in bounds add it
				if isWithinBounds(antinode1, grid) {
					antinodePositions[antinode1] = struct{}{}
				}
				if isWithinBounds(antinode2, grid) {
					antinodePositions[antinode2] = struct{}{}
				}
			}
		}
	}

	// Create Output grid, copy of og grid
	outputGrid := make([][]rune, len(grid))
	for i := range grid {
		outputGrid[i] = []rune(grid[i])
	}

	// Mark the antinode positions with #
	for pos := range antinodePositions {
		if isWithinBounds(pos, grid) {
			outputGrid[pos.y][pos.x] = '#'
		}
	}

	// Print the output grid with antinodes
	for _, row := range outputGrid {
		fmt.Println(string(row))
	}

	fmt.Printf("Unique locations: %d\n", len(antinodePositions))
}

// Check if point is within the grid bounds
func isWithinBounds(p Point, grid []string) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y])
}
