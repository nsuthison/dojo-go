package solutions

// Question: https://leetcode.com/problems/backspace-string-compare/
func backspaceCompare(S string, T string) bool {
	first := computeString(S)
	second := computeString(T)

	return first == second
}

func computeString(toCompute string) string {

	result := make([]rune, 0)

	for _, letter := range toCompute {
		if letter == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, letter)
		}
	}

	return string(result)
}
