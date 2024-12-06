package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5_part1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// ingesting data into maps and arrays
	rules := make(map[int][]int)
	pgNum := [][]int{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			section1 := strings.Split(line, "|")
			num1, err1 := strconv.Atoi(section1[0])
			num2, err2 := strconv.Atoi(section1[1])
			if err1 == nil && err2 == nil {
				// arr := []int{num1, num2}
				rules[num1] = append(rules[num1], num2)
			}
		}
		if strings.Contains(line, ",") {
			section2 := strings.Split(line, ",")
			arr := []int{}
			for _, num := range section2 {
				number, err := strconv.Atoi(num)
				if err == nil {
					arr = append(arr, number)
				}
			}
			pgNum = append(pgNum, arr)
		}

	}

	// fmt.Println(rules)
	// fmt.Println(pgNum)
	// solving puzzle
	output := 0

	for _, currentUpdate := range pgNum {
		isValid := true

		for i := len(currentUpdate) - 1; i >= 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
		}

		if isValid {
			output += currentUpdate[len(currentUpdate)/2]
		}
	}

	fmt.Println("Output for Day 5, Part 1:", output)

}

func day5_part2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// ingesting data into maps and arrays
	rules := make(map[int][]int)
	pgNum := [][]int{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			section1 := strings.Split(line, "|")
			num1, err1 := strconv.Atoi(section1[0])
			num2, err2 := strconv.Atoi(section1[1])
			if err1 == nil && err2 == nil {
				// arr := []int{num1, num2}
				rules[num1] = append(rules[num1], num2)
			}
		}
		if strings.Contains(line, ",") {
			section2 := strings.Split(line, ",")
			arr := []int{}
			for _, num := range section2 {
				number, err := strconv.Atoi(num)
				if err == nil {
					arr = append(arr, number)
				}
			}
			pgNum = append(pgNum, arr)
		}

	}

	// solving the puzzle
	output := 0

	for _, currentUpdate := range pgNum {
		isValid := true

		for i := len(currentUpdate) - 1; i >= 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
		}

		if !isValid {
			// need to iterate over the invalid updates
			// could do bogo sort for each value in invalid update and check if anything makes the list valid
			//
			output += currentUpdate[len(currentUpdate)/2]
		}
	}

	fmt.Println("Output for Day 5, Part 2:", output)

}

func helper_valid_updates(pgNum [][]int, rules map[int][]int) bool {
	for _, currentUpdate := range pgNum {
		isValid := true

		for i := len(currentUpdate) - 1; i >= 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
		}

		if isValid {
			return true
		}
	}

	return false
}
