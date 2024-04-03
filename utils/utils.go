package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"
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

func ParallelFor[T any](data []T, numGoroutines int, work WorkerFunc[T]) {
	var wg sync.WaitGroup
	chunkSize := (len(data) + numGoroutines - 1) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
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
