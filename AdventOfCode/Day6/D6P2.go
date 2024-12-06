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

	// Movement Deltas
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	// Func to simulate guard's movement
	isLoop := func(grid []string, startPos Point, startDir int) bool {
		visited := make(map[Point]int) // Track visits with direction
		pos := startPos
		dir := startDir
		steps := 0

		for steps < 1000 { // Step limit for loops
			state := Point{pos.x*4 + dir, pos.y} // encode pos and dir to state
			if visited[state] > 0 {
				return true // loop detected
				//Coming back to this this is the worst possible way to detect a loop what was I thinking
				//Goober guard could've walked thru where he started in a different direction and it wouldn't have been an infinite loop but I forgor to account for that
				//with this
			}
			visited[state]++

			nextPos := Point{pos.x + dx[dir], pos.y + dy[dir]}

			// OOB Checking
			if nextPos.y < 0 || nextPos.y >= len(grid) || nextPos.x < 0 || nextPos.x >= len(grid[0]) {
				return false // Guard exits the map
			}

			// Hitting an obstacle
			if grid[nextPos.y][nextPos.x] == '#' {
				// Turn right
				dir = (dir + 1) % 4
			} else {
				// Move forward
				pos = nextPos
			}
			steps++
		}
		return false
	}

	// Find all valid positions for obstructions
	validObstructionCount := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '.' && (x != guardPos.x || y != guardPos.y) {
				// Place obstruction temporarily
				grid[y] = grid[y][:x] + "#" + grid[y][x+1:]

				// Check if it causes a loop
				if isLoop(grid, guardPos, direction) {
					validObstructionCount++
				}

				// Remove the obstruction
				grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			}
		}
	}

	fmt.Println("Obstruction positions:", validObstructionCount)
}
