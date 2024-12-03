package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day3_part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	total := 0

	for _, line := range lines {

		matched, err := regexp.Compile(`mul\(\d+,\d+\)`)
		if err != nil {
			fmt.Println("Regex Error", err)
		}

		matches := matched.FindAllString(line, -1)
		for _, exp := range matches {
			if strings.HasPrefix(exp, "mul(") && strings.HasSuffix(exp, ")") {
				stringNum := exp[4 : len(exp)-1]
				sliceNum := strings.Split(stringNum, ",")
				if len(sliceNum) == 2 {
					leftNum, err1 := strconv.Atoi(sliceNum[0])
					rightNum, err2 := strconv.Atoi(sliceNum[1])
					if err1 == nil && err2 == nil {
						total += leftNum * rightNum
					}
				}
			}
		}
	}
	fmt.Println("Output for Day 3, Part 1:", total)
	return total
}

func day3_part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	total := 0
	do := true

	for _, line := range lines {

		matched, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		if err != nil {
			fmt.Println("Regex Error", err)
		}

		matches := matched.FindAllString(line, -1)
		for _, exp := range matches {
			if strings.HasPrefix(exp, "don't()") {
				do = false
			} else if strings.HasPrefix(exp, "do()") {
				do = true
			}
			if do && strings.HasPrefix(exp, "mul(") && strings.HasSuffix(exp, ")") {
				stringNum := exp[4 : len(exp)-1]
				sliceNum := strings.Split(stringNum, ",")
				if len(sliceNum) == 2 {
					leftNum, err1 := strconv.Atoi(sliceNum[0])
					rightNum, err2 := strconv.Atoi(sliceNum[1])
					if err1 == nil && err2 == nil {
						total += leftNum * rightNum
					}
				}
			}
		}
		// fmt.Println(matches)
	}
	fmt.Println("Output for Day 3, Part 2:", total)
	return total
}
