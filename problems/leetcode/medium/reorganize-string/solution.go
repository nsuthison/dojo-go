package solution

const notPossibleCase string = ""

// Question: https://leetcode.com/problems/reorganize-string/
func reorganizeString(S string) string {

	runes := []rune(S)

	for i := 0; i < len(runes) - 1; i++ {
		if string(runes[i]) == string(runes[i+1]) {

			isSwap := false
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

func categorizeNumberOfEachRuneIn(runes []rune) (letterMapper map[rune]int) {

	letterMapper = make(map[rune]int)

	for _, rune := range runes {

		if _, isExist := letterMapper[rune]; isExist {
			letterMapper[rune]++
		} else {
			letterMapper[rune] = 1
		}
	}

	return letterMapper
}
