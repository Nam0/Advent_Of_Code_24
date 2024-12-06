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
	fileName := "D6input.txt"
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

	// inital pos and dir
	var guardPos Point
	var direction int // 0 up 1 right 2 down 3 left
	found := false
	for y, row := range grid {
		for x, char := range row {
			if char == '^' {
				guardPos = Point{x, y}
				direction = 0
				found = true
				break
			} else if char == '>' {
				guardPos = Point{x, y}
				direction = 1
				found = true
				break
			} else if char == 'v' {
				guardPos = Point{x, y}
				direction = 2
				found = true
				break
			} else if char == '<' {
				guardPos = Point{x, y}
				direction = 3
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Deltas for movement
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	// tracked pos
	visited := make(map[Point]bool)
	visited[guardPos] = true

	// guard movement logic
	for {
		nextPos := Point{guardPos.x + dx[direction], guardPos.y + dy[direction]}

		// bounds check
		if nextPos.y < 0 || nextPos.y >= len(grid) || nextPos.x < 0 || nextPos.x >= len(grid[0]) {
			break
		}

		// if next pos is an obstacle
		if grid[nextPos.y][nextPos.x] == '#' {
			// turn right
			direction = (direction + 1) % 4
		} else {
			// go forward
			guardPos = nextPos
			visited[guardPos] = true
		}
	}

	fmt.Println("Pos visited:", len(visited))
}
