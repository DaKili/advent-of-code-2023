package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func ReadInput(inputFileName string) []string {
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(content), "\n")
}

type WorkerFunc[T any] func(subset []T)

func ParallelForStatic[T any](data []T, numGoroutines int, work WorkerFunc[T]) {
	if numGoroutines <= 0 {
		log.Println("ParallelForStatic: Cannot set number of Go routines to 0. Set to default: 1.")
		numGoroutines = 1
	}
	if numGoroutines == 16 && len(data) == 100 {
		fmt.Println("LAKSD")
	}
	var wg sync.WaitGroup
	chunkSize := len(data) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			work(data[start:end])
		}(start, end)
	}

	wg.Wait()
}

var MaxProcs = runtime.GOMAXPROCS(0)

func TestTask(t *testing.T, task func(string, int) int, inputFileName string, expected int) {
	testCases := []struct {
		parallelDegree int
	}{
		{parallelDegree: 1},
		{parallelDegree: 2},
		{parallelDegree: 4},
		{parallelDegree: 8},
		{parallelDegree: 16},
	}

	for _, testCase := range testCases {
		actual := task(inputFileName, testCase.parallelDegree)
		if actual != expected {
			t.Errorf("Test(%s, %d) = %d; want %d", inputFileName, testCase.parallelDegree, actual, expected)
		}
	}
}

func BenchmarkTask(b *testing.B, task func(string, int) int, inputFileName string) {
	for p := 1; p <= MaxProcs; p *= 2 {
		b.Run(fmt.Sprintf("%d_processes_of", p), func(b *testing.B) {
			for i := 1; i < b.N; i++ {
				task(inputFileName, p)
			}
		})
	}
	fmt.Println("")
}
