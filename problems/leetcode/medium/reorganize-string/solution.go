package solution

// Question: https://leetcode.com/problems/reorganize-string/
func reorganizeString(S string) string {

	runes := []rune(S)

	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == string(runes[i+1]) {
			runes[i], runes[i+1] = runes[i + 1], runes[i]
		}
	}

	return string(runes)
}

func swapRunes(runes *[]rune, firstIndex int, secondIndex int) {
	(*runes)[firstIndex], (*runes)[secondIndex] = (*runes)[secondIndex], (*runes)[firstIndex]
}
