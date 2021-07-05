package strquestions_test

import (
	"interviewq/strquestions"
	"testing"
)

func TestGetFirstNonRepeatingChar(t *testing.T) {
	scenarios := []struct {
		input    string
		expected rune
	}{
		{input: "aaabc", expected: 'b'},
		{input: "aaabcccbda", expected: 'd'},
		{input: "aaaabbbbccccdde", expected: 'e'},
		{input: "aaabbbccc", expected: '-'},
	}

	for _, s := range scenarios {
		got := strquestions.GetFirstNonRepeatingChar(s.input)
		if s.expected != got {
			t.Errorf("Did not meet the expected result. Expected '%q', got '%q'", s.expected, got)
		}
	}
}

func BenchmarkGetFirstNonRepeatingChar(b *testing.B) {
	input := "aaaabbbbccccdde"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		strquestions.GetFirstNonRepeatingChar(input)
	}
}

func TestGetFirstNonRepeatingCharNotOptimized(t *testing.T) {
	scenarios := []struct {
		input    string
		expected rune
	}{
		{input: "aaabc", expected: 'b'},
		{input: "aaabcccbda", expected: 'd'},
		{input: "aaaabbbbccccdde", expected: 'e'},
		{input: "aaabbbccc", expected: '-'},
	}

	for _, s := range scenarios {
		got := strquestions.GetFirstNonRepeatingCharNotOptimized(s.input)
		if s.expected != got {
			t.Errorf("Did not meet the expected result. Expected '%q', got '%q'", s.expected, got)
		}
	}
}

func BenchmarkGetFirstNonRepeatingCharNotOptimized(b *testing.B) {
	input := "aaaabbbbccccdde"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		strquestions.GetFirstNonRepeatingCharNotOptimized(input)
	}
}
