package solutions

func numDecodings(s string) int {
	return 0
}

//func calculateTail(tail []int, current) []int {
//
//}

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

