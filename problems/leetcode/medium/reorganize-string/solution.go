package solution

const notPossibleCase string = ""

// Question: https://leetcode.com/problems/reorganize-string/
func reorganizeString(S string) string {

	runes := []rune(S)

	for i := 0; i < len(runes) - 1; i++ {
		if string(runes[i]) == string(runes[i+1]) {

			isSwap := false

			// Case which need to swap tail to head to make reorganize valid
			if i + 2 == len(runes) && runes[0] != runes[i + 1]  {
				runes = append([]rune{runes[i + 1]}, runes[:i+1]...)
				return string(runes)
			}

			if i + 2 >= len(runes) {
				return notPossibleCase
			}

			for idxToSwap := i + 2; idxToSwap < len(runes); idxToSwap++ {
				if runes[i+1] != runes[idxToSwap] {
					swapRunes(&runes, i+1, idxToSwap)
					isSwap = true
					break
				}
			}

			if !isSwap {
				return notPossibleCase
			}
		}
	}

	return string(runes)
}

func swapRunes(runes *[]rune, firstIndex int, secondIndex int) {
	(*runes)[firstIndex], (*runes)[secondIndex] = (*runes)[secondIndex], (*runes)[firstIndex]
}
