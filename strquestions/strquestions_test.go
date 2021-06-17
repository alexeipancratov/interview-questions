package strquestions_test

import (
	"interviewq/strquestions"
	"testing"
)

func TestGetFirstRepeatingChar(t *testing.T) {
	scenarios := []struct {
		input    string
		expected rune
	}{
		{input: "aaabc", expected: 'b'},
		{input: "aaabcccbda", expected: 'd'},
		{input: "aaabbbccc", expected: '-'},
	}

	for _, s := range scenarios {
		got := strquestions.GetFirstNonRepeatingChar(s.input)
		if s.expected != got {
			t.Errorf("Did not meet the expected result. Expected '%q', got '%q'", s.expected, got)
		}
	}
}
