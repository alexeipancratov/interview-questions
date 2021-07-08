package memoization_test

import (
	"interviewq/dynamicprogramming/memoization"
	"testing"
)

func TestGridTraveler(t *testing.T) {
	scenarios := []struct {
		inputN   int
		inputM   int
		expected int
	}{
		{inputN: 0, inputM: 0, expected: 0},
		{inputN: 0, inputM: 1, expected: 0},
		{inputN: 1, inputM: 0, expected: 0},
		{inputN: 1, inputM: 1, expected: 1},
		{inputN: 2, inputM: 3, expected: 3},
		{inputN: 3, inputM: 2, expected: 3},
		{inputN: 3, inputM: 3, expected: 6},
		{inputN: 18, inputM: 18, expected: 2333606220},
	}

	for _, s := range scenarios {
		result := memoization.GridTraveler(s.inputN, s.inputM, nil)

		if result != s.expected {
			t.Errorf("Did not meet the expected result. Expected '%v', got '%v'", s.expected, result)
		}
	}
}

func BenchmarkGridTraveler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoization.GridTraveler(12, 12, nil)
	}
}

func BenchmarkGridTravelerNonOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoization.GridTravelerNonOptimized(12, 12)
	}
}
