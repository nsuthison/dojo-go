package solutions

import "strconv"

// Question: https://leetcode.com/problems/decode-ways/
func numDecodings(s string) int {

	previousTailInfo := make(map[int]int, 0)

	for idx, runeDigit := range s {

		if !isAbleToDecode(idx, s) {
			return 0
		}

		currentDigit, _ := strconv.Atoi(string(runeDigit))
		previousTailInfo = nextTailInfo(previousTailInfo, currentDigit)
	}

	var result int
	for _, count := range previousTailInfo {
		result += count
	}

	return result
}

func isAbleToDecode(currentIdx int, input string) bool {
	currentDigit, _ := strconv.Atoi(string(input[currentIdx]))

	if currentIdx == 0 && currentDigit == 0 {
		return false
	}

	if currentIdx > 0 {
		previousDigit, _ := strconv.Atoi(string(input[currentIdx-1]))

		if (previousDigit == 0 || previousDigit > 2) && currentDigit == 0 {
			return false
		}
	}

	return true
}

func nextTailInfo(previousTailInfos map[int]int, toAppend int) (nextTailInfo map[int]int) {

	nextTailInfo = make(map[int]int, 0)

	if len(previousTailInfos) == 0 {
		nextTailInfo[toAppend] = 1
	}

	for number, count := range previousTailInfos {

		if toAppend != 0 {
			nextTailInfo[toAppend] = nextTailInfo[toAppend] + count
		}

		if isBeAbleToCombine(number, toAppend) {

			combine := createCombine(number, toAppend)
			nextTailInfo[combine] = nextTailInfo[combine] + count
		}
	}

	return nextTailInfo
}

func createCombine(tens int, ones int) int {
	return (tens * 10) + ones
}

func isBeAbleToCombine(head int, tail int) bool {

	switch head {
	case 1:
		return true
	case 2:
		if tail <= 6 {
			return true
		}
	}

	return false
}
