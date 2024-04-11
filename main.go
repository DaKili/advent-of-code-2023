package main

import (
	"fmt"

	day_01 "github.com/dakili/advent-of-code-2023/day-01"
	day_02 "github.com/dakili/advent-of-code-2023/day-02"
	day_03 "github.com/dakili/advent-of-code-2023/day-03"
)

func main() {
	fmt.Printf("Day01 Task1 solution: %v\n", day_01.Task1("day-01/input.txt", 8))
	fmt.Printf("Day01 Task2 solution: %v\n", day_01.Task2("day-01/input.txt", 8))
	fmt.Printf("Day02 Task1 solution: %v\n", day_02.Task1("day-02/input.txt", 8))
	fmt.Printf("Day02 Task2 solution: %v\n", day_02.Task2("day-02/input.txt", 8))
	fmt.Printf("Day03 Task1 solution: %v\n", day_03.Task1("day-03/input.txt", 8))
	fmt.Printf("Day03 Task2 solution: %v\n", day_03.Task2("day-03/input.txt", 8))
}
