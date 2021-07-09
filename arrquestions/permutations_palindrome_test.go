package arrquestions_test

import (
	"interviewq/arrquestions"
	"testing"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	scenarios := []struct {
		input    string
		expected bool
	}{
		{input: "level", expected: true},
		{input: "on meoln no loenm", expected: true},
		{input: "waas tii tca aws", expected: true},
		{input: "hnnaha", expected: true},
		{input: "elizabeth", expected: false},
		{input: "becky", expected: false},
	}

	for _, scenario := range scenarios {
		result := arrquestions.IsPermutationOfPalindrome(scenario.input)
		if result != scenario.expected {
			t.Fatalf("Expected to return '%v' for '%s' string, but got '%v'", scenario.expected, scenario.input, result)
		}
	}
}
