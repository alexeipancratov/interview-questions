package arrquestions_test

import (
	"interviewq/arrquestions"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	scenarios := []struct {
		input    []int
		item     int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			item:     5,
			expected: 4,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			item:     10,
			expected: -1,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			item:     1,
			expected: 0,
		},
	}

	for _, s := range scenarios {
		result := arrquestions.BinarySortSearch(&s.input, s.item)
		if result != s.expected {
			t.Fatalf("Expected %d but got %d", s.expected, result)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		arrquestions.BinarySortSearch(&input, 5)
	}
}

func TestLinearSearch(t *testing.T) {
	scenarios := []struct {
		input    []int
		item     int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			item:     5,
			expected: 4,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			item:     10,
			expected: -1,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			item:     1,
			expected: 0,
		},
	}

	for _, s := range scenarios {
		result := arrquestions.LinearSearch(&s.input, s.item)
		if result != s.expected {
			t.Fatalf("Expected %d but got %d", s.expected, result)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		arrquestions.LinearSearch(&input, 5)
	}
}
