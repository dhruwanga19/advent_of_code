package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day7_part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		res, _ := strconv.Atoi(parts[0])
		numbers := strings.Split(parts[1], " ")
		integers := []int{}
		for _, num := range numbers {
			number, _ := strconv.Atoi(num)
			integers = append(integers, number)
		}
		result[res] = integers
	}

	output := 0

	for res, numbers := range result {
		if check_combinations_p1(numbers, res) {
			output += res
		}

	}

	fmt.Println("Output for Day 7, Part 1:", output)

	return output
}

func day7_part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		res, _ := strconv.Atoi(parts[0])
		numbers := strings.Split(parts[1], " ")
		integers := []int{}
		for _, num := range numbers {
			number, _ := strconv.Atoi(num)
			integers = append(integers, number)
		}
		result[res] = integers
	}

	output := 0

	for res, numbers := range result {
		if check_combinations_p2(numbers, res) {
			output += res
		}

	}

	fmt.Println("Output for Day 7, Part 2:", output)

	return output
}

func check_combinations_p1(numbers []int, result int) bool {
	n := len(numbers)
	if n == 0 {
		return false
	}
	if n == 1 {
		return numbers[0] == result
	}

	var recursive func(int, int) bool
	recursive = func(idx int, current int) bool {
		if idx == n {
			return current == result
		}
		return recursive(idx+1, current+numbers[idx]) || recursive(idx+1, current*numbers[idx])
	}

	return recursive(1, numbers[0])
}

func check_combinations_p2(numbers []int, result int) bool {
	n := len(numbers)
	if n == 0 {
		return false
	}
	if n == 1 {
		return numbers[0] == result
	}

	var recursive func(int, int) bool
	recursive = func(idx int, current int) bool {
		if idx == n {
			return current == result
		}
		concat := current*int(math.Pow10(int(math.Log10(float64(numbers[idx]))+1))) + numbers[idx]
		return recursive(idx+1, current+numbers[idx]) || recursive(idx+1, current*numbers[idx]) || recursive(idx+1, concat)
	}

	return recursive(1, numbers[0])
}
