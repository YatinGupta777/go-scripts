package main

import (
	"testing"
)

func simulateCacheClearance(b *testing.B) {
	// Simulate some warm-up or setup
	data := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		// Simulate cache clearing by accessing a large array
		for j := range data {
			_ = data[j]
		}
	}
}

func BenchmarkFunction1(b *testing.B) {
	simulateCacheClearance(b)
	// Reset the timer to exclude setup time
	b.ResetTimer()

	input := [2]int64{0, 0}
	for i := 0; i < b.N; i++ {
		function1(input)
	}
}

func BenchmarkFunction2(b *testing.B) {
	simulateCacheClearance(b)
	// Reset the timer to exclude setup time
	b.ResetTimer()

	input := [2]int64{0, 0}
	for i := 0; i < b.N; i++ {
		function2(input)
	}
}
