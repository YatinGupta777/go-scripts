package main

import (
	"fmt"
	"time"
)

const increment_times = 1_000_000

// This script shows how instruction level parallelism optimizes code
// Both functions perform the same operation but one is seemingly faster

// Function 1: Simulate a simple operation
func function1(s [2]int64) [2]int64 {
	for i := 0; i < increment_times; i++ {
		s[0]++
		if s[0]%2 == 0 {
			s[1]++
		}
	}
	return s
}

// Function 2: Has more instructions which can be parallized
func function2(s [2]int64) [2]int64 {
	for i := 0; i < increment_times; i++ {
		v := s[0]
		s[0] = v + 1
		if v%2 != 0 {
			s[1]++
		}
	}
	return s
}

func measureExecutionTime(fn func([2]int64) [2]int64, input [2]int64, name string) time.Duration {
	start := time.Now()
	fn(input)
	duration := time.Since(start)
	fmt.Printf("%s executed in %v\n", name, duration)
	return duration
}

func main() {
	// Measure the execution time of both functions
	input := [2]int64{1, 2}

	fmt.Println("Comparing execution times:")
	measureExecutionTime(function1, input, "Function1")
	measureExecutionTime(function2, input, "Function2")
}
