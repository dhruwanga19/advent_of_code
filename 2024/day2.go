package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2_part1(input string) int {

	inputReport := strings.Split(strings.TrimSpace(input), "\n")

	var reports [][]int
	for _, line := range inputReport {
		report := strings.Fields(line)
		var level []int
		for _, numStr := range report {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			level = append(level, num)
		}
		reports = append(reports, level)
	}

	safe := 0
	for _, report := range reports {
		// fmt.Print(report)
		isIncreasing := true
		isSafe := true
		for i := 1; i < len(report); i++ {
			diff := report[i] - report[i-1]
			if diff == 0 {
				isSafe = false
				break
			}
			if diff > 3 || diff < -3 {
				isSafe = false
				break
			}

			if i == 1 && diff < 0 {
				isIncreasing = false

			} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
				isSafe = false
				break
			}

		}
		if isSafe {
			safe += 1
		}

	}

	fmt.Println("Output for day 2 part 1:", safe)

	return safe
}

func day2_part2(input string) int {
	inputReport := strings.Split(strings.TrimSpace(input), "\n")

	var reports [][]int
	for _, line := range inputReport {
		report := strings.Fields(line)
		var level []int
		for _, numStr := range report {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			level = append(level, num)
		}
		reports = append(reports, level)
	}

	safe := 0

	for _, report := range reports {
		isSafe := false
		isSafe = reportIsSafe(report)

		if !isSafe {
			for i := 0; i < len(report); i++ {
				newReport := []int{}
				newReport = append(newReport, report[:i]...)
				newReport = append(newReport, report[i+1:]...)
				isSafe = reportIsSafe(newReport)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			safe++
		}

	}

	fmt.Println("Output for day 2 part 2:", safe)
	return safe
}

// helper function to check if report is safe
func reportIsSafe(level []int) bool {
	isIncreasing := true
	isSafe := true
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]
		if diff == 0 {
			isSafe = false
			break
		}
		if diff > 3 || diff < -3 {
			isSafe = false
			break
		}

		if i == 1 && diff < 0 {
			isIncreasing = false

		} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			isSafe = false
			break
		}
	}
	return isSafe
}
