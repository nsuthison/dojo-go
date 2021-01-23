package solutions

// Questiom: https://leetcode.com/problems/expressive-words/
func expressiveWords(S string, words []string) int {
	return 0
}

func createStretchyMap(word string) []StretchyMap {
	previousLetter := rune(word[0])
	repeatLetterCount := 0

	stretchMap := make([]StretchyMap, 0)

	for idx := 1; idx <= len(word); idx++ {
		if idx < len(word) && rune(word[idx]) == previousLetter {
			repeatLetterCount++
		} else {

			isLetterStretchy := isLetterStretchy(repeatLetterCount)

			switch repeatLetterCount {
			case 2:
				stretchMap = append(stretchMap, StretchyMap{previousLetter, isLetterStretchy})
				stretchMap = append(stretchMap, StretchyMap{previousLetter, isLetterStretchy})
			default:
				stretchMap = append(stretchMap, StretchyMap{previousLetter, isLetterStretchy})
			}

			repeatLetterCount = 1
		}

		previousLetter = rune(word[idx])
	}

	return stretchMap
}

func isLetterStretchy(repeatLetterCount int) bool {
	return repeatLetterCount >= 3
}

// StretchyMap Map to represent that is the letter stretchy
type StretchyMap struct {
	Letter     rune
	IsStretchy bool
}
