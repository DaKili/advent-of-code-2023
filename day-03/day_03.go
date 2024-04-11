package day_03

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/dakili/advent-of-code-2023/utils"
)

type enginePart struct {
	start int
	end   int
	value int
}

type enginePartLine struct {
	parts       []enginePart
	symbolIndex []int
	gearIndex   []int
}

func Task1(inputFileName string, parallelDeree int) int {
	// NEED TO PARALLELIZE
	input := utils.ReadInput(inputFileName)
	if input == nil {
		fmt.Printf("Could not read %s\n", inputFileName)
		return 0
	}

	enginePartLines := createEnginePartLines(input)

	sum := getEnginePartSum(enginePartLines, len(input))

	return sum
}

func Task2(inputFileName string, parallelDeree int) int {
	// NEED TO PARALLELIZE
	input := utils.ReadInput(inputFileName)
	if input == nil {
		fmt.Printf("Could not read %s\n", inputFileName)
		return 0
	}

	enginePartLines := createEnginePartLines(input)

	sum := getGearEnginePartSum(enginePartLines, len(input))

	return sum
}

func getEnginePartSum(enginePartLines []enginePartLine, inputLength int) int {
	sum := 0
	for i, line := range enginePartLines {
		if i == 0 {
			for _, part := range line.parts {
				if HasPartSymbolInRange(part, enginePartLines[i]) ||
					HasPartSymbolInRange(part, enginePartLines[i+1]) {
					sum += part.value
				}
			}
		} else if i < inputLength-1 {
			for _, part := range line.parts {
				if HasPartSymbolInRange(part, enginePartLines[i-1]) ||
					HasPartSymbolInRange(part, enginePartLines[i]) ||
					HasPartSymbolInRange(part, enginePartLines[i+1]) {
					sum += part.value
				}
			}

		} else {
			for _, part := range line.parts {
				if HasPartSymbolInRange(part, enginePartLines[i-1]) ||
					HasPartSymbolInRange(part, enginePartLines[i]) {
					sum += part.value
				}
			}
		}
	}

	return sum
}

func getGearEnginePartSum(enginePartLines []enginePartLine, inputLength int) int {
	sum := 0
	for i, line := range enginePartLines {
		if i == 0 {
			for _, gearPosition := range line.gearIndex {
				adjacentParts := []enginePart{}
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i])...)
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i+1])...)
				if len(adjacentParts) == 2 {
					sum += adjacentParts[0].value * adjacentParts[1].value
				}
			}

		} else if i < inputLength-1 {
			for _, gearPosition := range line.gearIndex {
				adjacentParts := []enginePart{}
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i-1])...)
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i])...)
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i+1])...)
				if len(adjacentParts) == 2 {
					sum += adjacentParts[0].value * adjacentParts[1].value
				}
			}
		} else {
			for _, gearPosition := range line.gearIndex {
				adjacentParts := []enginePart{}
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i-1])...)
				adjacentParts = append(adjacentParts, getAdjacentParts(gearPosition, enginePartLines[i])...)
				if len(adjacentParts) == 2 {
					sum += adjacentParts[0].value * adjacentParts[1].value
				}
			}
		}
	}

	return sum
}

func getAdjacentParts(gearPosition int, epl enginePartLine) []enginePart {
	adjacentParts := []enginePart{}
	start := gearPosition - 1
	end := gearPosition + 1
	for _, part := range epl.parts {
		if part.end >= start && part.start <= end {
			adjacentParts = append(adjacentParts, part)
		}
	}
	return adjacentParts
}

func HasPartSymbolInRange(part enginePart, belowLine enginePartLine) bool {
	for _, symbolIndex := range belowLine.symbolIndex {
		if part.start-1 <= symbolIndex && part.end+1 >= symbolIndex {
			return true
		}
	}
	return false
}

func createEnginePartLines(input []string) []enginePartLine {
	enginePartLines := []enginePartLine{}
	for _, line := range input {
		epl := enginePartLine{}
		for i := 0; i < len(line); i++ {
			if rune(line[i]) == '.' {
				continue
			} else if unicode.IsDigit(rune(line[i])) {
				newPart := enginePart{start: i}

				if i+1 < len(line) && unicode.IsDigit(rune(line[i+1])) {
					if i+2 < len(line) && unicode.IsDigit(rune(line[i+2])) {
						newPart.end = i + 2
						newPart.value, _ = strconv.Atoi(line[newPart.start : newPart.end+1])
						i += 2
					} else {
						newPart.end = i + 1
						newPart.value, _ = strconv.Atoi(line[newPart.start : newPart.end+1])
						i += 1
					}
				} else {
					newPart.end = i
					newPart.value, _ = strconv.Atoi(line[newPart.start : newPart.end+1])
				}

				epl.parts = append(epl.parts, newPart)
			} else if line[i] == '*' {
				epl.gearIndex = append(epl.gearIndex, i)
				epl.symbolIndex = append(epl.symbolIndex, i)
			} else {
				epl.symbolIndex = append(epl.symbolIndex, i)
			}
		}
		enginePartLines = append(enginePartLines, epl)
	}
	return enginePartLines
}
