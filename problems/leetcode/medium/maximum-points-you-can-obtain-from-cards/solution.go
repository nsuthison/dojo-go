package solution

import . "math"

// Question: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
func maxScore(cardPoints []int, k int) int {

	numberOfCardsToTakeFromFront := k
	numberOfCardsToTakeFromBack := 0

	summationMemoizeFromFront := make([]int, 0)
	summationMemoizeFromBack := make([]int, 0)

	maxScore := 0

	for numberOfCardsToTakeFromBack <= k {

		sumFromFront := getSummationOfCardsFrom(Front, cardPoints, numberOfCardsToTakeFromFront, &summationMemoizeFromFront)
		sumFromBack := getSummationOfCardsFrom(Back, cardPoints, numberOfCardsToTakeFromBack, &summationMemoizeFromBack)
		currentSum := sumFromFront + sumFromBack

		if currentSum > maxScore {
			maxScore = currentSum
		}

		numberOfCardsToTakeFromFront--
		numberOfCardsToTakeFromBack++
	}

	return maxScore
}

func getSummationOfCardsFrom(direction LinearDirection, cardPoints []int, numberOfCardsToTake int, memoize *[]int) (sumResult int) {

	if numberOfCardsToTake <= 0 {
		return 0
	}

	pointerIndexStartingPoint := getPointerIndexStartingPointFrom(direction, len(cardPoints))

	memoizeSize := len(*memoize)

	if memoizeSize < numberOfCardsToTake {
		for i := memoizeSize; i < numberOfCardsToTake; i++ {

			if i == 0 {
				*memoize = append(*memoize, cardPoints[pointerIndexStartingPoint])
				continue
			}

			lastSummationInMemoize := (*memoize)[i-1]
			nextCardValueToCalculate := cardPoints[int(Abs(float64(pointerIndexStartingPoint-i)))]

			*memoize = append(*memoize, lastSummationInMemoize+nextCardValueToCalculate)
		}
	}

	return (*memoize)[numberOfCardsToTake-1]
}

func getPointerIndexStartingPointFrom(direction LinearDirection, arraySize int) int {
	var pointerIndexStartingPoint int

	switch direction {
	case Front:
		pointerIndexStartingPoint = 0
	case Back:
		pointerIndexStartingPoint = arraySize - 1
	}

	return pointerIndexStartingPoint
}
