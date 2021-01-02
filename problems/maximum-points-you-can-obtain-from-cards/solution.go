package solution

// Question: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

func maxScore(cardPoints []int, k int) int {
	return 0
}

func createSumMapFromFront(cards []int, numberOfCardsToSelect int) []int {
	sumCardValueFromFront := make([]int, numberOfCardsToSelect)

	sumCardValueFromFront[0] = cards[0]
	for i := 1; i < numberOfCardsToSelect; i++ {
		sumCardValueFromFront[i] = sumCardValueFromFront[i - 1] + cards[i]
	}

	return sumCardValueFromFront
}

