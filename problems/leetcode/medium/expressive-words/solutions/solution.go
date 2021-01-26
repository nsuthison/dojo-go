package solutions

import models "github.com/nsuthison/dojo-go/problems/leetcode/medium/expressive-words/models"

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

func createLetterInfosFrom(word string) []models.LetterInfo {

	letterInfos := make([]models.LetterInfo, 0)
	repeatLetterCount := 1

	for idx := 1; idx <= len(word); idx++ {
		previousLetter := rune(word[idx-1])

		if idx < len(word) && rune(word[idx]) == previousLetter {
			repeatLetterCount++
		} else {

			letterInfos = append(letterInfos, models.LetterInfo{Letter: previousLetter, LetterCount: repeatLetterCount})

			repeatLetterCount = 1
		}
	}

	return letterInfos
}

func isWordStretchy(word string, letterInfos []models.LetterInfo) bool {

	idx := 0

	for _, letterInfo := range letterInfos {

		// If idx equal or more than word length it should already check all letterInfos.
		// Or else, word have more letter than the base one
		if idx >= len(word) {
			return false
		}

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
	}

	// If letterInfos is exhaust but still has some letter in word to check, this word is shorter than the base one.
	if idx < len(word) {
		return false
	}

	return true
}

func isLetterStretchy(repeatLetterCount int) bool {
	return repeatLetterCount >= 3
}
