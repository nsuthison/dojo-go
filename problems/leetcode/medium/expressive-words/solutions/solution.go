package solutions

// Questiom: https://leetcode.com/problems/expressive-words/
func expressiveWords(S string, words []string) int {

	if len(S) == 0 {
		return 0
	}

	validWordCount := 0
	letterInfos := createLetterInfosFrom(S)

	for _, word := range words {
		if isWordStretchy(word, letterInfos) {
			validWordCount++
		}
	}

	return validWordCount
}

func createLetterInfosFrom(word string) []LetterInfo {
	previousLetter := rune(word[0])
	repeatLetterCount := 1

	letterInfos := make([]LetterInfo, 0)

	for idx := 1; idx <= len(word); idx++ {
		if idx < len(word) && rune(word[idx]) == previousLetter {
			repeatLetterCount++
		} else {

			letterInfos = append(letterInfos, LetterInfo{previousLetter, repeatLetterCount})

			repeatLetterCount = 1
		}

		if idx < len(word) {
			previousLetter = rune(word[idx])
		}
	}

	return letterInfos
}

func isLetterStretchy(repeatLetterCount int) bool {
	return repeatLetterCount >= 3
}

func isWordStretchy(word string, letterInfos []LetterInfo) bool {

	idx := 0

	for _, letterInfo := range letterInfos {
		if idx < len(word) {

			if letterInfo.Letter != rune(word[idx]) {
				return false
			}

			repeatLetterCount := 0
			for idx < len(word) && letterInfo.Letter == rune(word[idx]) {
				idx++
				repeatLetterCount++
			}

			if isLetterStretchy(letterInfo.LetterCount) {
				if letterInfo.LetterCount < repeatLetterCount {
					return false
				}
			} else {
				if letterInfo.LetterCount != repeatLetterCount {
					return false
				}
			}
		} else {
			return false
		}
	}

	if idx < len(word) {
		return false
	}

	return true
}

// LetterInfo Map to represent that is the letter stretchy
type LetterInfo struct {
	Letter      rune
	LetterCount int
}
