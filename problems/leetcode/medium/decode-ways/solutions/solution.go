package solutions

import "strconv"

//func numDecodings(s string) int {
//
//	tails := make([]int, 0)
//	hasNumber := false
//	isLegit := true
//
//	for _, runeDigit := range s {
//
//		digit, _ := strconv.Atoi(string(runeDigit))
//
//		if !hasNumber && digit == 0 {
//			continue
//		}
//
//		hasNumber = true
//
//		tails, isLegit = nextTails(tails, digit)
//		if !isLegit {
//			return 0
//		}
//	}
//
//	return len(tails)
//}

func numDecodings(s string) int {

	previousTailInfos := make(map[int]int, 0)

	for idx, runeDigit := range s {

		digit, _ := strconv.Atoi(string(runeDigit))

		if idx == 0 && digit == 0 {
			return 0
		}

		if idx > 0 {
			previousDigit, _ := strconv.Atoi(string(s[idx - 1]))

			if (previousDigit == 0 || previousDigit > 2) && digit == 0 {
				return 0
			}
		}

		previousTailInfos = nextTailInfo(previousTailInfos, digit)
	}

	var result int
	for _, count := range previousTailInfos {
		result += count
	}

	return result
}

func nextTailInfo(previousTailInfos map[int]int, toAppend int) (nextTailInfo map[int]int) {

	nextTailInfo = make(map[int]int, 0)

	if len(previousTailInfos) == 0 {
		nextTailInfo[toAppend] = 1
	}

	for number, count := range previousTailInfos {

		if toAppend == 0 {
			if isBeAbleToCombine(number, toAppend) {
				combine := (number * 10) + toAppend

				nextTailInfo[combine] = nextTailInfo[combine] + count
			} else {
				continue
			}
		} else {
			nextTailInfo[toAppend] = nextTailInfo[toAppend] + count

			if isBeAbleToCombine(number, toAppend) {
				combine := (number * 10) + toAppend

				nextTailInfo[combine] = nextTailInfo[combine] + count
			}
		}
	}

	return nextTailInfo
}

func nextTails(tails []int, toAppend int) (result []int, isAbleToGet bool) {

	result = make([]int, 0)

	if len(tails) == 0 {
		return append(result, toAppend), true
	}

	if toAppend == 0 {
		for _, tail := range tails {
			if !isBeAbleToCombine(tail, 0) {
				continue
			} else {
				combine := (tail * 10) + toAppend
				result = append(result, combine)
			}
		}
	} else {

		for _, tail := range tails {
			result = append(result, toAppend)

			if isBeAbleToCombine(tail, toAppend) {
				combine := (tail * 10) + toAppend

				result = append(result, combine)
			}
		}
	}

	if len(result) == 0 {
		return nil, false
	} else {
		return result, true
	}
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