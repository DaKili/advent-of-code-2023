package day_01

import (
	"testing"

	"github.com/dakili/advent-of-code-2023/utils"
)

var inputFileName = "input.txt"

func TestDay01Task1(t *testing.T) {
	const expected = 55017
	utils.TestTask(t, Task1, inputFileName, expected)
}

func TestDay01Task2(t *testing.T) {
	const expected = 53539
	utils.TestTask(t, Task2, inputFileName, expected)
}

func BenchmarkDay01Task1(b *testing.B) {
	utils.BenchmarkTask(b, Task1, inputFileName)
}

func BenchmarkDay01Task2(b *testing.B) {
	utils.BenchmarkTask(b, Task2, inputFileName)
}
