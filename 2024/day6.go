package main

import (
	"fmt"
	"strings"
)

func day6_part1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	output := 0

	startRow := -1
	startCol := -1
	startDirection := 0

	for row := range grid {
		if startRow >= 0 {
			break
		}
		for col := range grid[row] {
			if grid[row][col] == '^' {
				startRow = row
				startCol = col
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visLocs := make(map[[2]int]bool)

	for {
		visLocs[[2]int{startRow, startCol}] = true

		currentDirection := directions[startDirection]
		nextRow := startRow + currentDirection[0]
		nextCol := startCol + currentDirection[1]

		if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
			break
		}

		if grid[nextRow][nextCol] == '#' {
			startDirection = (startDirection + 1) % 4
			currentDirection = directions[startDirection]
			nextRow = startRow + currentDirection[0]
			nextCol = startCol + currentDirection[1]

			if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
				break
			}
		}
		startRow = nextRow
		startCol = nextCol
	}
	output = len(visLocs)

	fmt.Println("Output Day 6 Part 1:", output)
}

func day6_part2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	// output := 0

	startRow := -1
	startCol := -1
	startDirection := 0

	for row := range grid {
		if startRow >= 0 {
			break
		}
		for col := range grid[row] {
			if grid[row][col] == '^' {
				startRow = row
				startCol = col
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visLocs := make(map[[2]int]bool)

	for {
		visLocs[[2]int{startRow, startCol}] = true

		currentDirection := directions[startDirection]
		nextRow := startRow + currentDirection[0]
		nextCol := startCol + currentDirection[1]

		if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
			break
		}

		if grid[nextRow][nextCol] == '#' {
			startDirection = (startDirection + 1) % 4
			currentDirection = directions[startDirection]
			nextRow = startRow + currentDirection[0]
			nextCol = startCol + currentDirection[1]

			if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
				break
			}
		}
		startRow = nextRow
		startCol = nextCol
	}
	// output = len(visLocs)
}
