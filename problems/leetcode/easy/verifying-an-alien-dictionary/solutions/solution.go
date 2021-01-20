package solutions

// Question: https://leetcode.com/problems/verifying-an-alien-dictionary/
func isAlienSorted(words []string, order string) bool {
	letterOrderMappingRule := createLetterOrderMappingRuleFrom(order)

	for i := 0; i < len(words)-1; i++ {
		isLeftWordLesserThanRight := isLeftWordLesserThanRight(words[i], words[i+1], letterOrderMappingRule)

		if !isLeftWordLesserThanRight {
			return false
		}
	}

	return true
}

func isLeftWordLesserThanRight(leftWord string, rightWord string, orderRule map[rune]int) bool {

	leftWordAsRunes := []rune(leftWord)
	rightWordAsRunes := []rune(rightWord)

	for i := 0; i < len(leftWordAsRunes); i++ {

		if i >= len(rightWordAsRunes) {
			return false
		}

		leftLetterOrderPoint := orderRule[leftWordAsRunes[i]]
		rightLetterOrderPoint := orderRule[rightWordAsRunes[i]]

		switch {
		case leftLetterOrderPoint < rightLetterOrderPoint:
			return true
		case leftLetterOrderPoint > rightLetterOrderPoint:
			return false
		default:
			continue
		}
	}

	return true
}

func createLetterOrderMappingRuleFrom(order string) map[rune]int {

	letterOrderMapping := make(map[rune]int)
	for idx, letter := range order {
		letterOrderMapping[letter] = idx
	}

	return letterOrderMapping
}
