package strquestions

func GetFirstNonRepeatingChar(str string) rune {
	occurrences := [26]int{}

	for i := 0; i < len(str); i++ {
		currentChar := rune(str[i])
		occurrences[currentChar-'a']++
	}

	for i, o := range occurrences {
		if o == 1 {
			return rune(i + 'a')
		}
	}

	return '-'
}
