package main

import (
	"fmt"
	"time"

	day_01 "github.com/dakili/advent-of-code-2023/day-01"
	day_02 "github.com/dakili/advent-of-code-2023/day-02"
	day_03 "github.com/dakili/advent-of-code-2023/day-03"
)

func main() {
	// call the task you want to run day_02.Task1()
	printBenchmarkTemplate()
	benchmark(day_01.Task1)
	benchmark(day_01.Task2)
	benchmark(day_02.Task1)
	benchmark(day_02.Task2)
	benchmark(day_03.Task1)
	benchmark(day_03.Task2)
}

func printBenchmarkTemplate() {
	fmt.Printf("%15s %15s %15s %15s %15s\n", "Name", "Solution", "Best", "Worst", "Average")
	fmt.Printf("_______________________________________________________________________________\n")
}

func benchmark(day_task func() (string, int)) {
	const repetitions = 10
	var sum time.Duration
	var min, max = time.Hour, time.Nanosecond
	name := ""
	solution := 0
	for i := 0; i < repetitions; i++ {
		start := time.Now()
		name, solution = day_task()
		executionTime := time.Since(start)
		if executionTime < min {
			min = executionTime
		} else if executionTime > max {
			max = executionTime
		}
		sum += executionTime
	}
	avg := sum / time.Duration(repetitions)
	fmt.Printf("%13s | %15d %13vms %13vms %13vms\n", name, solution, min.Microseconds(), max.Microseconds(), avg.Microseconds())
}
