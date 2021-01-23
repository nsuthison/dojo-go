package solutions

// Questiom: https://leetcode.com/problems/expressive-words/
func expressiveWords(S string, words []string) int {

	validWordCount := 0
	stretchyInfo := createStretchyInfosFrom(S)

	for _, word := range words {
		if isWordStretchy(word, stretchyInfo) {
			validWordCount++
		}
	}

	return validWordCount
}

func createStretchyInfosFrom(word string) []StretchyInfo {
	previousLetter := rune(word[0])
	repeatLetterCount := 0

	stretchyInfos := make([]StretchyInfo, 0)

	for idx := 1; idx <= len(word); idx++ {
		if idx < len(word) && rune(word[idx]) == previousLetter {
			repeatLetterCount++
		} else {

			isLetterStretchy := isLetterStretchy(repeatLetterCount)

			switch repeatLetterCount {
			case 2:
				stretchyInfos = append(stretchyInfos, StretchyInfo{previousLetter, isLetterStretchy})
				stretchyInfos = append(stretchyInfos, StretchyInfo{previousLetter, isLetterStretchy})
			default:
				stretchyInfos = append(stretchyInfos, StretchyInfo{previousLetter, isLetterStretchy})
			}

			repeatLetterCount = 1
		}

		if idx < len(word) {
			previousLetter = rune(word[idx])
		}
	}

	return stretchyInfos
}

func isLetterStretchy(repeatLetterCount int) bool {
	return repeatLetterCount >= 3
}

func isWordStretchy(word string, StretchyInfos []StretchyInfo) bool {

	idx := 0

	for _, stretchyInfo := range StretchyInfos {
		if idx < len(word) {

			if stretchyInfo.Letter != rune(word[idx]) {
				return false
			}

			if stretchyInfo.IsStretchy {
				for idx < len(word) && stretchyInfo.Letter == rune(word[idx]) {
					idx++
				}
			} else {
				idx++
			}
		} else {
			return false
		}
	}

	return true
}

// StretchyInfo Map to represent that is the letter stretchy
type StretchyInfo struct {
	Letter     rune
	IsStretchy bool
}
