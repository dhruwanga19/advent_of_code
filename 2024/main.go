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

func main() {

	day := 3
	input := getInputFile(day)

	// day 1
	// startTime := time.Now()
	// day1_part1(input)
	// day1_part2(input)
	// elapsedTime := time.Since(startTime)
	// fmt.Println("Operation took:", elapsedTime)

	//day 2
	// startTime := time.Now()
	// day2_part1(input)
	// day2_part2(input)
	// elapsedTime := time.Since(startTime)
	// fmt.Println("Operation took:", elapsedTime)

	//day 3
	startTime := time.Now()
	day3_part1(input)
	day3_part2(input)
	elapsedTime := time.Since(startTime)
	fmt.Println("Operation took:", elapsedTime)
}
