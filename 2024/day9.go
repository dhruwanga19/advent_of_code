package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day9_part1(input string) int {
	line := strings.TrimSpace(input)
	block := []rune{}
	isBlock := true
	fileID := 0

	for _, numLine := range line {
		num, _ := strconv.Atoi(string(numLine))
		if isBlock {
			for i := 0; i < num; i++ {
				block = append(block, rune('0'+fileID))
			}
			fileID++
		} else {
			for i := 0; i < num; i++ {
				block = append(block, '.')
			}
		}
		isBlock = !isBlock
	}

	// fmt.Println(block)

	empty := 0
	lastFile := len(block) - 1

	for block[empty] != '.' {
		empty++
	}
	for block[lastFile] == '.' {
		lastFile--
	}

	for empty <= lastFile {
		block[empty] = block[lastFile]
		block[lastFile] = '.'

		for block[empty] != '.' {
			empty++
		}

		for block[lastFile] == '.' {
			lastFile--
		}
	}

	blockId := 0
	output := 0

	for block[blockId] != '.' {
		output += blockId * int(block[blockId]-'0')
		blockId++
	}

	fmt.Println("Output Day 9, Part 1:", output)

	// output := 0
	return output
}

func day9_part2(input string) int {
	line := strings.TrimSpace(input)
	block := []rune{}
	isBlock := true
	fileID := 0

	for _, numLine := range line {
		num, _ := strconv.Atoi(string(numLine))
		if isBlock {
			for i := 0; i < num; i++ {
				block = append(block, rune('0'+fileID))
			}
			fileID++
		} else {
			for i := 0; i < num; i++ {
				block = append(block, '.')
			}
		}
		isBlock = !isBlock
	}

	// fmt.Println(block)

	for currentFile := fileID - 1; currentFile >= 0; currentFile-- {
		fileBlocks := []int{}

		for i, block := range block {
			if block == rune('0'+currentFile) {
				fileBlocks = append(fileBlocks, i)
			}
		}

		freeStart := -1
		freeLength := 0

		for i := 0; i < fileBlocks[0]; i++ {
			if block[i] != '.' {
				freeStart = -1
				freeLength = 0
				continue
			}

			if freeStart == -1 {
				freeStart = i
			}

			freeLength++

			if freeLength == len(fileBlocks) {
				break
			}
		}

		if freeLength == len(fileBlocks) {
			for i := 0; i < freeLength; i++ {
				block[fileBlocks[i]] = '.'
				block[freeStart+i] = rune('0' + currentFile)
			}
		}
	}
	// blockId := 0
	output := 0

	for blockId := 0; blockId < len(block); blockId++ {
		if block[blockId] != '.' {

			output += blockId * int(block[blockId]-'0')
		}
		// blockId++
	}

	fmt.Println("Output Day 9, Part 2:", output)

	// output := 0
	return output
}
