package memoization_test

import (
	"interviewq/dynamicprogramming/memoization"
	"testing"
)

func TestCanSum(t *testing.T) {
	scenarios := []struct {
		inputN     int
		inputItems []int
		expected   bool
	}{
		{inputN: 7, inputItems: []int{2, 3}, expected: true},
		{inputN: 7, inputItems: []int{5, 3, 4, 7}, expected: true},
		{inputN: 7, inputItems: []int{2, 4}, expected: false},
		{inputN: 8, inputItems: []int{2, 3, 5}, expected: true},
		{inputN: 300, inputItems: []int{7, 14}, expected: false},
	}

	for _, s := range scenarios {
		result := memoization.CanSum(s.inputN, s.inputItems, nil)

		if result != s.expected {
			t.Errorf("Did not meet the expected result. Expected '%v', got '%v'", s.expected, result)
		}
	}
}

func BenchmarkCanSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoization.CanSum(300, []int{7, 14}, nil)
	}
}

func BenchmarkCanSumNonOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoization.CanSumNonOptimized(300, []int{7, 14})
	}
}
