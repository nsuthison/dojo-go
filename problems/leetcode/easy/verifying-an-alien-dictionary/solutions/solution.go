package solutions

func isAlienSorted(words []string, order string) bool {
	return false
}

func createLetterOrderMappingFrom(order string) map[rune]int {

	letterOrderMapping := make(map[rune]int)
	for idx, letter := range order {
		letterOrderMapping[letter] = idx
	}

	return letterOrderMapping
}