package main

import (
	"fmt"
	"strings"
)

func day4_part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	/* word search allows words to be
	horizontal,
	vertical,
	diagonal,
	written backwards,
	or even overlapping other words
	*/

	word := "XMAS"

	output := 0

	directions := [][2]int{
		{0, 1},   // right
		{0, -1},  // left
		{-1, 0},  // up
		{1, 0},   // down
		{-1, 1},  // up-right
		{-1, -1}, // up-left
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if search_word(grid, word, i, j, dir[0], dir[1]) {
					output++
				}
			}
		}

	}

	fmt.Println("Output for Day 4, Part 1:", output)

	return output
}

func search_word(grid [][]rune, word string, row int, col int, rowOffset int, colOffset int) bool {
	for k := 0; k < len(word); k++ {
		nx, ny := row+k*rowOffset, col+k*colOffset
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != rune(word[k]) {
			return false
		}
	}
	return true
}

func day4_part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, row := range lines {
		grid[i] = []rune(row)
	}

	output := 0

	directions := [][2][2]int{
		{
			{-1, -1}, //up-left
			{1, 1},   //down-right
		},
		{
			{-1, 1}, //up-right
			{1, -1}, //down-left
		},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {

			if grid[row][col] != 'A' {
				continue
			}

			isMas := true

			for _, dir := range directions {
				row1Index := row + dir[0][0]
				col1Index := col + dir[0][1]

				row2Index := row + dir[1][0]
				col2Index := col + dir[1][1]

				if row1Index < 0 || row1Index >= len(grid) || col1Index < 0 || col1Index >= len(grid[row]) ||
					row2Index < 0 || row2Index >= len(grid) || col2Index < 0 || col2Index >= len(grid[row]) {
					isMas = false
					break
				}

				if (grid[row1Index][col1Index] == 'M' && grid[row2Index][col2Index] == 'S') ||
					(grid[row1Index][col1Index] == 'S' && grid[row2Index][col2Index] == 'M') {
					continue
				}

				isMas = false
				break

			}
			if isMas {
				output++
			}

		}
	}

	fmt.Println("Output for Day 4, Part 2:", output)
	return output
}
