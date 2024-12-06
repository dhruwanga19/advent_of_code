package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func getInputFile(day int) string {

	inputDirectory := "inputs"
	inputFile := filepath.Join(inputDirectory, fmt.Sprintf("inputday%d.txt", day))

	if _, err := os.Stat(inputFile); err == nil {
		content, err := os.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("Could not read file %s", inputFile)
		}
		return string(content)
	} else if os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", inputFile)
	} else {
		log.Fatalf("Error checking file %s: %v", inputFile, err)
	}

	return inputFile
}
func runDay(day int) {
	input := getInputFile(day)
	startTime := time.Now()
	switch day {
	case 1:
		day1_part1(input)
		day1_part2(input)
	case 2:
		day2_part1(input)
		day2_part2(input)
	case 3:
		day3_part1(input)
		day3_part2(input)
	case 4:
		day4_part1(input)
		day4_part2(input)
	case 5:
		day5_part1(input)
		day5_part2(input)
	default:
		fmt.Printf("Day %d is not implemented\n", day)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Day %d operation took: %s\n", day, elapsedTime)
}

func main() {
	day := 5
	runDay(day)
}
