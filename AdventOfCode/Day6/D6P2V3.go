package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
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

	// Func for guard movement
	moveGuard := func(tempGrid []string, pos Point, dir int) (Point, int) {
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
				return pos, dir
			}
		}
	}

	totalTests := 0
	infiniteCount := 0
	var mu sync.Mutex

	// Worker function for Multi threading
	worker := func(start, end int, results chan<- struct{ tests, loops int }, wg *sync.WaitGroup) {
		defer wg.Done()
		tests := 0
		loops := 0
		for i := start; i < end; i++ {
			y := i / len(grid[0])
			x := i % len(grid[0])
			if grid[y][x] != '.' {
				continue
			}

			tempGrid := make([]string, len(grid))
			copy(tempGrid, grid)
			tempGrid[y] = tempGrid[y][:x] + "#" + tempGrid[y][x+1:]

			tempGuardPos := guardPos
			tempDirection := direction

			positionHistory := make(map[string]bool)
			for {
				state := fmt.Sprintf("%d,%d,%d", tempGuardPos.x, tempGuardPos.y, tempDirection)
				if positionHistory[state] {
					loops++
					break
				}
				positionHistory[state] = true

				newPos, newDir := moveGuard(tempGrid, tempGuardPos, tempDirection)
				if newPos == tempGuardPos && newDir == tempDirection {
					break
				}
				tempGuardPos, tempDirection = newPos, newDir
			}
			tests++
		}
		results <- struct{ tests, loops int }{tests, loops}
	}

	// Concurrency setup stuff
	numWorkers := 8
	numCells := len(grid) * len(grid[0])
	chunkSize := (numCells + numWorkers - 1) / numWorkers
	results := make(chan struct{ tests, loops int }, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > numCells {
			end = numCells
		}
		wg.Add(1)
		go worker(start, end, results, &wg)
	}

	// Collection
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		mu.Lock()
		totalTests += result.tests
		infiniteCount += result.loops
		mu.Unlock()
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Time To Execute: %v\n", duration) // 11 seconds could increase num of workers but idk if my CPU can cope
	fmt.Printf("Total Tests: %d\n", totalTests)
	fmt.Printf("Infinite Loops: %d\n", infiniteCount)
}
