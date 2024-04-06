package day_02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dakili/advent-of-code-2023/utils"
)

var inputFileName = "day-02/input.txt"

func Task1() {
	defer utils.Timer("Day-02 Task1")()
	const (
		redMax   = 12
		greenMax = 13
		blueMax  = 14
	)

	input := utils.ReadInput(inputFileName)
	if input == nil {
		fmt.Printf("Could not read %s\n", inputFileName)
		return
	}

	possibleGames := 0
	// Parallelization has barely any improvement due to overhead.
	utils.ParallelForStatic(input, 4, func(subset []string) {
		possibleGames += getPossibleGamesSum(subset, redMax, greenMax, blueMax)
	})

	fmt.Printf("Solution: %v\n", possibleGames)
}

func getPossibleGamesSum(lines []string, redMax, greenMax, blueMax int) int {
	gameIdSum := 0
	for _, line := range lines {
		gameId, rest := cutGameId(line)
		draws := strings.Split(rest, ";")
		if isGamePossible(draws, redMax, greenMax, blueMax) {
			gameIdSum += gameId
		}
	}
	return gameIdSum
}

func isGamePossible(draws []string, redMax, greenMax, blueMax int) bool {
	// Quicker than regex
	for _, draw := range draws {
		draw = strings.Trim(draw, " ")
		colors := strings.Split(draw, ",")
		for _, color := range colors {
			color := strings.Trim(color, " ")
			colorSlice := strings.Split(color, " ")
			count, _ := strconv.Atoi(colorSlice[0])
			color = colorSlice[1]
			switch color {
			case "red":
				if count > redMax {
					return false
				}
			case "green":
				if count > greenMax {
					return false
				}
			case "blue":
				if count > blueMax {
					return false
				}
			}
		}
	}
	return true
}

func cutGameId(line string) (int, string) {
	gameSplit := strings.Split(line, ":")
	rest, _ := strings.CutPrefix(gameSplit[0], "Game ")
	convertedId, _ := strconv.Atoi(rest)
	return convertedId, strings.Trim(gameSplit[1], " ") // This will always pass for advent of code.
}
