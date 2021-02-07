package solutions

func fullJustify(words []string, maxWidth int) []string {
	return make([]string, 0)
}

func groupTextForEachLine(words []string, maxWidth int) [][]string {
	toReturn := make([][]string, 0)

	eachLine := make([]string, 0)
	for _, word := range words {
		totalLen(eachLine)
	}
}

func totalLen(words []string) int {
	totalLen := 0

	for _, word := range words {
		totalLen += len(word)
	}

	return totalLen
}
