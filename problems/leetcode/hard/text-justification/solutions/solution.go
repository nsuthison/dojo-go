package solutions

// Question: https://leetcode.com/problems/text-justification/
func fullJustify(words []string, maxWidth int) []string {
	return make([]string, 0)
}

// func groupTextForEachLine(words []string, maxWidth int) [][]string {
// 	toReturn := make([][]string, 0)

// 	eachLine := make([]string, 0)
// 	for _, word := range words {
// 		totalLen(eachLine)
// 	}
// }

func totalLenWithSpace(words []string) int {
	totalLen := 0

	for _, word := range words {
		totalLen += len(word)
	}

	return totalLen + totalSpaceFor(len(words))
}

func totalSpaceFor(numberOfWord int) int {
	if numberOfWord <= 1 {
		return 0
	}

	return numberOfWord - 1
}
