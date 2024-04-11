package day_02

import (
	"testing"

	"github.com/dakili/advent-of-code-2023/utils"
)

var inputFileName = "input.txt"

func TestDay02Task1(t *testing.T) {
	const expected = 2776
	utils.TestTask(t, Task1, inputFileName, expected)
}

func TestDay02Task2(t *testing.T) {
	const expected = 68638
	utils.TestTask(t, Task2, inputFileName, expected)
}

func BenchmarkDay02Task1(b *testing.B) {
	utils.BenchmarkTask(b, Task1, inputFileName)
}

func BenchmarkDay02Task2(b *testing.B) {
	utils.BenchmarkTask(b, Task2, inputFileName)
}
