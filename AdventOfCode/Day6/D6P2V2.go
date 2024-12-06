package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	startTime := time.Now()
	fileName := "test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Map slices
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

	// Track visited posisions
	visited := make(map[Point]bool)
	visited[guardPos] = true

	// Func for guard movement
	moveGuard := func(tempGrid []string, visited map[Point]bool, pos Point, dir int) (Point, int) {
		for {
			nextPos := Point{pos.x + dx[dir], pos.y + dy[dir]}

			// Bounds checking stuff
			if nextPos.y < 0 || nextPos.y >= len(tempGrid) || nextPos.x < 0 || nextPos.x >= len(tempGrid[0]) {
				return pos, dir
			}

			// Hitting an obstacle
			if tempGrid[nextPos.y][nextPos.x] == '#' {
				// Turn right
				dir = (dir + 1) % 4
			} else {
				// Move forward
				pos = nextPos
				visited[pos] = true
				return pos, dir
			}
		}
	}

	// Total valid obstruction tests
	totalTests := 0
	infiniteCount := 0

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				continue
			}

			// temp grid with obstacle at curr poss
			tempGrid := make([]string, len(grid))
			copy(tempGrid, grid)
			tempGrid[y] = tempGrid[y][:x] + "#" + tempGrid[y][x+1:]

			// rest guard pos and dir
			tempVisited := make(map[Point]bool)
			tempGuardPos := guardPos
			tempDirection := direction
			tempVisited[tempGuardPos] = true

			// guard movement
			positionHistory := make(map[string]bool)
			for {
				state := fmt.Sprintf("%d,%d,%d", tempGuardPos.x, tempGuardPos.y, tempDirection)
				if positionHistory[state] {
					infiniteCount++
					break
				}
				positionHistory[state] = true

				newPos, newDir := moveGuard(tempGrid, tempVisited, tempGuardPos, tempDirection)
				if newPos == tempGuardPos && newDir == tempDirection {
					// Guard cant move
					break
				}
				tempGuardPos, tempDirection = newPos, newDir
			}

			totalTests++
		}
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Time To Execute: %v\n", duration) // optomized to aprox 44 seconds yucky
	fmt.Printf("Total Tests: %d\n", totalTests)
	fmt.Printf("Infinite Loops: %d\n", infiniteCount)
}
