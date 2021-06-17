package strquestions

func GetFirstNonRepeatingChar(str string) rune {
	occurrenceCounts := [26]int{}

	for i := 0; i < len(str); i++ {
		currentChar := rune(str[i])
		occurrenceCounts[currentChar-'a']++
	}

	for i, oc := range occurrenceCounts {
		if oc == 1 {
			return rune(i + 'a')
		}
	}

	return '-'
}

func GetFirstNonRepeatingCharNotOptimized(str string) rune {
	for i := 0; i < len(str); i++ {
		var repeatingChar rune = '-'

		for j := 0; j < len(str); j++ {
			if i != j && str[i] == str[j] {
				repeatingChar = rune(str[i])
			}
		}

		if repeatingChar == '-' {
			return rune(str[i])
		}
	}
	return '-'
}
