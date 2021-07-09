package arrquestions

import "regexp"

func IsPermutationOfPalindrome(str string) bool {
	charsRepetitions := make(map[rune]int)
	for i := 0; i < len(str); i++ {
		if isLetter(str[i]) {
			charsRepetitions[rune(str[i])]++
		}
	}

	oneOddIsFound := false
	for _, frequency := range charsRepetitions {
		if frequency%2 == 1 && oneOddIsFound {
			return false
		}
		oneOddIsFound = frequency%2 == 1
	}

	return true
}

func isLetter(char byte) bool {
	matches, _ := regexp.MatchString("[a-z]", string(char))

	return matches
}
