package main

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
