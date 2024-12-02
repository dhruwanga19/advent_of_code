package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	day := 1
	input := getInputFile(day)
	day1_part1(input)
	day1_part2(input)
}
