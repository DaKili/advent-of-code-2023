package day_01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dakili/advent-of-code-2023/utils"
)

var InputFileName = "day-01/input.txt"
var Parallel_parts = 8

func Task1() {
	defer utils.Timer("Day-01 Task1")()

	input := utils.ReadInput(InputFileName)
	if input == nil {
		fmt.Printf("Could not read %s\n", InputFileName)
		return
	}

	sum := 0
	dictionary := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}
	utils.ParallelForStatic(input, Parallel_parts, func(subset []string) {
		sum += processLines(subset, dictionary)
	})

	fmt.Printf("Solution: %v\n", sum)
}

func Task2() {
	defer utils.Timer("Day-01 Task2")()

	input := utils.ReadInput(InputFileName)
	if input == nil {
		fmt.Printf("Could not read %s\n", InputFileName)
		return
	}

	var dictionary = map[string]string{
		"1": "1", "one": "1",
		"2": "2", "two": "2",
		"3": "3", "three": "3",
		"4": "4", "four": "4",
		"5": "5", "five": "5",
		"6": "6", "six": "6",
		"7": "7", "seven": "7",
		"8": "8", "eight": "8",
		"9": "9", "nine": "9",
	}

	sum := 0
	utils.ParallelForStatic(input, Parallel_parts, func(subset []string) {
		sum += processLines(subset, dictionary)
	})
	fmt.Printf("Solution: %v\n", sum)
}

func processLines(lines []string, dictionary map[string]string) int {
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		first, last := getFirstAndLastDigit(line, dictionary)
		concatenated := dictionary[first] + dictionary[last]
		calibrationValue, err := strconv.Atoi(concatenated)
		if err == nil {
			sum += calibrationValue
		}
	}
	return sum
}

func getFirstAndLastDigit(line string, dictionary map[string]string) (string, string) {
	firstIndex := -1
	firstDigit := "0"
	for k, v := range dictionary {
		index := strings.Index(line, k)
		if index != -1 {
			if firstIndex == -1 {
				firstIndex = index
				firstDigit = dictionary[v]

			} else if firstIndex > index {
				firstIndex = index
				firstDigit = dictionary[v]
			}
		}
	}

	lastIndex := -1
	lastDigit := "0"
	for k, v := range dictionary {
		index := strings.LastIndex(line, k)
		if index != -1 && index > lastIndex {
			lastIndex = index
			lastDigit = dictionary[v]
		}
	}

	return firstDigit, lastDigit
}
