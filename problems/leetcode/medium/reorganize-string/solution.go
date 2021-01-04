package solution

// Question: https://leetcode.com/problems/reorganize-string/
func reorganizeString(S string) string {

	runes := []rune(S)

	for i := 0; i < len(runes) - 1; i++ {
		if string(runes[i]) == string(runes[i+1]) {
			isSwap := false

			if i + 2 >= len(runes) {
				return ""
			}

			for idxToSwap := i + 2; idxToSwap < len(runes); idxToSwap++ {
				if runes[i+1] != runes[idxToSwap] {
					swapRunes(&runes, i+1, idxToSwap)
					isSwap = true
					break
				}
			}

			if !isSwap {
				return ""
			}
		}
	}

	return string(runes)
}

func swapRunes(runes *[]rune, firstIndex int, secondIndex int) {
	(*runes)[firstIndex], (*runes)[secondIndex] = (*runes)[secondIndex], (*runes)[firstIndex]
}
