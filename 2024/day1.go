package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func day1_part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left_list []int
	var right_list []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		leftNum, err1 := strconv.Atoi(numbers[0])
		rightNum, err2 := strconv.Atoi(numbers[1])
		if err1 == nil && err2 == nil {
			left_list = append(left_list, leftNum)
			right_list = append(right_list, rightNum)
		}
	}

	sort.Ints(left_list)
	sort.Ints(right_list)

	total_distance := 0

	for i := 0; i < len(left_list); i++ {
		total_distance += int(math.Abs(float64(right_list[i] - left_list[i])))
	}

	fmt.Println("Output Day 1, Part 1: Total distance is:", total_distance)
	return total_distance
}

func day1_part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left_list []int
	var right_list []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		leftNum, err1 := strconv.Atoi(numbers[0])
		rightNum, err2 := strconv.Atoi(numbers[1])
		if err1 == nil && err2 == nil {
			left_list = append(left_list, leftNum)
			right_list = append(right_list, rightNum)
		}
	}

	right_map := make(map[int]int)

	for _, num := range right_list {
		right_map[num]++
	}

	similarity_score := 0

	for _, num := range left_list {
		similarity_score += num * right_map[num]
	}

	fmt.Println("Output Day 1, Part 2: Similarity score is:", similarity_score)
	return similarity_score
}
