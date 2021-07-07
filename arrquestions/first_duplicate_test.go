package arrquestions_test

import (
	"interviewq/arrquestions"
	"testing"
)

var scenarios = []struct {
	input    []int
	expected int
}{
	{input: []int{1, 2, 4, 2, 1}, expected: 2},
	{input: []int{1, 2, 4, 1, 2}, expected: 1},
	{input: []int{1, 2, 3, 4}, expected: -1},
}

func TestGetFirstDuplicate(t *testing.T) {
	for _, scenario := range scenarios {
		result := arrquestions.GetFirstDuplicate(&scenario.input)
		if result != scenario.expected {
			t.Fatalf("Expected the first duplicate to be %v, but got %v", scenario.expected, result)
		}
	}
}

func BenchmarkGetFirstDuplicateNSquare(b *testing.B) {
	input := []int{1, 2, 4, 2, 1}
	for i := 0; i < b.N; i++ {
		arrquestions.GetFirstDuplicateNSquare(&input)
	}
}

func BenchmarkGetFirstDuplicateNonOptimizedMap(b *testing.B) {
	input := []int{1, 2, 4, 2, 1}
	for i := 0; i < b.N; i++ {
		arrquestions.GetFirstDuplicateNonOptimizedMap(&input)
	}
}

func BenchmarkGetFirstDuplicateMap(b *testing.B) {
	input := []int{1, 2, 4, 2, 1}
	for i := 0; i < b.N; i++ {
		arrquestions.GetFirstDuplicateMap(&input)
	}
}

func BenchmarkGetFirstDuplicate(b *testing.B) {
	input := []int{1, 2, 4, 2, 1}
	for i := 0; i < b.N; i++ {
		arrquestions.GetFirstDuplicate(&input)
	}
}
