package strquestions

func GetFirstNonRepeatingChar(str string) rune {
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
