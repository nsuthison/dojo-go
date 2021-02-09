package solutions

// Question: https://leetcode.com/problems/text-justification/
func fullJustify(words []string, maxWidth int) []string {

	wordsSeparateToEachLine := groupWordsToLineFrom(words, maxWidth)

	return fillInSpaceFor(wordsSeparateToEachLine, maxWidth)
}

func groupWordsToLineFrom(words []string, maxWidth int) [][]string {
	toReturn := make([][]string, 0)

	eachLine := make([]string, 0)
	for _, word := range words {
		if totalLenWithSpace(eachLine)+1+len(word) > maxWidth {
			toReturn = append(toReturn, eachLine)
			eachLine = make([]string, 0)
		}

		eachLine = append(eachLine, word)
	}

	if len(eachLine) > 0 {
		toReturn = append(toReturn, eachLine)
	}

	return toReturn
}

func fillInSpaceFor(wordSeparateToLine [][]string, maxWidth int) []string {

	justifyTexts := make([]string, 0)

	for _, wordsInEachLine := range wordSeparateToLine {
		numberOfLetterInLine := totalLetterIn(wordsInEachLine)

		totalSpaceChar := maxWidth - numberOfLetterInLine
		numberOfSpaceGap := getNumberOfSpaceGapInTextFor(wordsInEachLine)

		spaceWidthBetweenWord := totalSpaceChar / numberOfSpaceGap
		spaceFraction := totalSpaceChar % numberOfSpaceGap

		justifyText := createJustifyTextFrom(wordsInEachLine, spaceWidthBetweenWord, spaceFraction)

		justifyTexts = append(justifyTexts, justifyText)
	}

	return justifyTexts
}

func getNumberOfSpaceGapInTextFor(words []string) int {

	numberOfWords := len(words)

	if numberOfWords == 1 {
		return 1
	}

	return numberOfWords - 1
}

func createJustifyTextFrom(words []string, spaceWidthBetweenWord int, spaceFraction int) string {
	toReturn := ""

	for idx, word := range words {

		spaceToFillIn := spaceWidthBetweenWord

		switch idx {
		case 0:
			spaceToFillIn = spaceWidthBetweenWord + spaceFraction
		case len(words) - 1:
			spaceToFillIn = 0
		default:
			spaceToFillIn = spaceWidthBetweenWord
		}

		toReturn += word
		toReturn += createSpaces(spaceToFillIn)
	}

	return toReturn
}

func createSpaces(size int) string {
	toReturn := ""

	for i := 0; i < size; i++ {
		toReturn += " "
	}

	return toReturn
}

func totalLenWithSpace(words []string) int {
	return totalLetterIn(words) + totalSpaceFor(len(words))
}

func totalLetterIn(words []string) int {
	totalLen := 0

	for _, word := range words {
		totalLen += len(word)
	}

	return totalLen
}

func totalSpaceFor(numberOfWord int) int {
	if numberOfWord <= 1 {
		return 0
	}

	return numberOfWord - 1
}
