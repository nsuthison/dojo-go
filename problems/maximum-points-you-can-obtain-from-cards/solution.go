package solution

// Question: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
import . "math"

func maxScore(cardPoints []int, k int) int {
	//
	//numberOfCardsToSelect := k
	//sumCardValueFromFront := createSumMapFromFront(cardPoints, numberOfCardsToSelect)
	//
	//numberOfCardsToSelectFromFront := k
	//numberOfCardsToSelectFromBack := 0
	//
	//for false {
	//
	//}

	return 0
}

func createSumMapFromFront(cards []int, numberOfCardsToSelect int) []int {
	sumCardValueFromFront := make([]int, numberOfCardsToSelect)

	sumCardValueFromFront[0] = cards[0]
	for i := 1; i < numberOfCardsToSelect; i++ {
		sumCardValueFromFront[i] = sumCardValueFromFront[i-1] + cards[i]
	}

	return sumCardValueFromFront
}

func getSummationOfCardsFrom(direction LinearDirection, cardPoints []int, numberOfCardsToTake int, memoize *[]int) (sumResult int) {

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
