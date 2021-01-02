package solution

// Question: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

func maxScore(cardPoints []int, k int) int {

	numberOfCardsToSelect := k
	sumCardValueFromFront := createSumMapFromFront(cardPoints, numberOfCardsToSelect)

	cardPointsIndexFromLast := len(cardPoints) - 1

	sumCardValueFromBack := make([]int, 0)
	sumCardValueFromBack = append(sumCardValueFromBack, cardPoints[cardPointsIndexFromLast])

	numberOfCardTakenFromFront := numberOfCardsToSelect - 1
	numberOfCardTakenFromBack := 1

	maxScore := sumCardValueFromFront[numberOfCardTakenFromFront - 1]

	for numberOfCardTakenFromBack <= numberOfCardsToSelect {
		currentSummation := sumCardValueFromFront[numberOfCardTakenFromFront - 1] + sumCardValueFromBack[numberOfCardTakenFromBack - 1]

		if maxScore < currentSummation {
			maxScore = currentSummation
		}
	}

	//for numberOfCardTakenFromBack <= numberOfCardsToSelect {
	//	currentSumScore := sumCardValueFromFront[numberOfCardTakenFromFront - 1] + cardPoints[cardPointsLastIndex - ]
	//}

	return maxScore
}

func createSumMapFromFront(cards []int, numberOfCardsToSelect int) []int {
	sumCardValueFromFront := make([]int, numberOfCardsToSelect)

	sumCardValueFromFront[0] = cards[0]
	for i := 1; i < numberOfCardsToSelect; i++ {
		sumCardValueFromFront[i] = sumCardValueFromFront[i - 1] + cards[i]
	}

	return sumCardValueFromFront
}

func getSummationOfCardsFromFront(cardPoints []int, numberOfCardsToTake int, memoize *[]int) (sumResult int) {

	memoizeSize := len(*memoize)

	if memoizeSize < numberOfCardsToTake {
		for i := memoizeSize; i < numberOfCardsToTake; i++ {

			if i == 0 {
				*memoize = append(*memoize, cardPoints[0])
				continue
			}

			lastSummationInMemoize := (*memoize)[i-1]
			nextCardValueToCalculate := cardPoints[i]

			*memoize = append(*memoize, lastSummationInMemoize + nextCardValueToCalculate)
		}
	}

	return (*memoize)[numberOfCardsToTake - 1]
}

func getSummationOfCardsFromBack(cardPoints []int, numberOfCardsToTake int, memoize *[]int) (sumResult int) {
	memoizeSize := len(*memoize)
	lastCardPointsIndex := len(cardPoints) - 1

	if memoizeSize < numberOfCardsToTake {
		for i := memoizeSize; i < numberOfCardsToTake; i++ {

			if i == 0 {
				*memoize = append(*memoize, cardPoints[lastCardPointsIndex])
				continue
			}

			lastSummationInMemoize := (*memoize)[i-1]
			nextCardValueToCalculate := cardPoints[lastCardPointsIndex - i]

			*memoize = append(*memoize, lastSummationInMemoize + nextCardValueToCalculate)
		}
	}

	return (*memoize)[numberOfCardsToTake - 1]
}

