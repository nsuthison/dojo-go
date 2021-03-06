package solutions

// Question: https://leetcode.com/problems/text-justification/
func fullJustify(words []string, maxWidth int) []string {

	wordsSeparateToEachLine := groupWordsToLineFrom(words, maxWidth)

	return fillInSpaceFor(wordsSeparateToEachLine, maxWidth)
}

func groupWordsToLineFrom(words []string, maxWidth int) [][]string {
	toReturn := make([][]string, 0)

	eachLine := []string{words[0]}

	for idx := 1; idx < len(words); idx++ {
		if totalLetterWithSpace(eachLine)+len(" ")+len(words[idx]) > maxWidth {
			toReturn = append(toReturn, eachLine)

			eachLine = make([]string, 0)
		}

		eachLine = append(eachLine, words[idx])
	}

	if len(eachLine) > 0 {
		toReturn = append(toReturn, eachLine)
	}

	return toReturn
}

func fillInSpaceFor(wordSeparateToLine [][]string, maxWidth int) []string {

	justifyTexts := make([]string, 0)
	lastLineIdx := len(wordSeparateToLine) - 1

	for idx, wordsInEachLine := range wordSeparateToLine {

		justifyText := ""
		if idx == lastLineIdx {
			justifyText = createJustifyTextForLastLineFrom(wordsInEachLine, maxWidth)
		} else {
			justifyText = createJustifyTextFrom(wordsInEachLine, maxWidth)
		}

		justifyTexts = append(justifyTexts, justifyText)
	}

	return justifyTexts
}

func createJustifyTextFrom(words []string, maxWidth int) string {

	totalSpaceChar := maxWidth - totalLetterIn(words)
	numberOfSpaceGap := getNumberOfSpaceGapInTextFor(words)

	spaceWidthBetweenWord := totalSpaceChar / numberOfSpaceGap
	spaceFraction := totalSpaceChar % numberOfSpaceGap

	toReturn := ""
	for idx, word := range words {

		spaceToFillIn := 0

		if len(words) == 1 {
			spaceToFillIn = spaceWidthBetweenWord
		} else {
			switch i := idx; {
			case i < spaceFraction:
				spaceToFillIn = spaceWidthBetweenWord + 1
			case i == len(words)-1:
				spaceToFillIn = 0
			default:
				spaceToFillIn = spaceWidthBetweenWord
			}
		}

		toReturn += word
		toReturn += createSpaces(spaceToFillIn)
	}

	return toReturn
}

func getNumberOfSpaceGapInTextFor(words []string) int {

	numberOfWords := len(words)

	if numberOfWords == 1 {
		return 1
	}

	return numberOfWords - 1
}

func createJustifyTextForLastLineFrom(words []string, maxWidth int) string {
	justifyText := ""
	for wordIdx, word := range words {
		lastWordIdx := len(words) - 1

		if wordIdx == lastWordIdx {
			justifyText += word
			justifyText += createSpaces(maxWidth - len(justifyText))
		} else {
			justifyText += word
			justifyText += createSpaces(1)
		}
	}

	return justifyText
}

func createSpaces(size int) string {
	toReturn := ""

	for i := 0; i < size; i++ {
		toReturn += " "
	}

	return toReturn
}

func totalLetterWithSpace(words []string) int {
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
