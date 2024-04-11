package day_03

import (
	"testing"

	"github.com/dakili/advent-of-code-2023/utils"
)

var inputFileName = "input.txt"

func TestDay03Task1(t *testing.T) {
	const expected = 532428
	utils.TestTask(t, Task1, inputFileName, expected)
}

func TestDay03Task2(t *testing.T) {
	const expected = 84051670
	utils.TestTask(t, Task2, inputFileName, expected)
}

func BenchmarkDay03Task1(b *testing.B) {
	utils.BenchmarkTask(b, Task1, inputFileName)
}

func BenchmarkDay03Task2(b *testing.B) {
	utils.BenchmarkTask(b, Task2, inputFileName)
}
