package solutions

import (
	"github.com/nsuthison/dojo-go/problems/leetcode/medium/reorganize-string/models"
	"sort"
)

const notPossibleCase string = ""

// Question: https://leetcode.com/problems/reorganize-string/

/// Solution is that,
/// Desc sort letter in string by its count number
/// Then use 2 pointer from top to bottom to append letter to create result
/// It might have left over letter from the last sorted letter array
/// Insert left over to the result from the start of the result string by jumping each time by 2
func reorganizeString(S string) string {

	runes := []rune(S)

	if len(S) <= 1 {
		return S
	}

	letterMapper := categorizeNumberOfEachRuneIn(runes)
	letterInfos := createLetterInStringInfosFrom(letterMapper)
	sortedLetterInfos := descSort(letterInfos)

	if len(sortedLetterInfos) <= 1 {
		return notPossibleCase
	}

	currentResult := make([]rune, 0)

	firstIdxPointer := 0
	secondIdxPointer := 1

	canSelectNextPointer := true

	for canSelectNextPointer {
		canSelectNextPointer = false

		for haveLetterToInsertAt(firstIdxPointer, sortedLetterInfos) &&
			haveLetterToInsertAt(secondIdxPointer, sortedLetterInfos) {

			currentResult = append(currentResult, sortedLetterInfos[firstIdxPointer].Letter)
			currentResult = append(currentResult, sortedLetterInfos[secondIdxPointer].Letter)

			sortedLetterInfos[firstIdxPointer].NumberOfLetter--
			sortedLetterInfos[secondIdxPointer].NumberOfLetter--
		}

		if !haveLetterToInsertAt(firstIdxPointer, sortedLetterInfos) {
			var nextSelectedPointer int
			nextSelectedPointer, canSelectNextPointer = trySelectNextPointerFor(firstIdxPointer, sortedLetterInfos, secondIdxPointer)

			if canSelectNextPointer {
				firstIdxPointer = nextSelectedPointer
			}
		}

		if !haveLetterToInsertAt(secondIdxPointer, sortedLetterInfos) {
			var nextSelectedPointer int
			nextSelectedPointer, canSelectNextPointer = trySelectNextPointerFor(secondIdxPointer, sortedLetterInfos, firstIdxPointer)

			if canSelectNextPointer {
				secondIdxPointer = nextSelectedPointer
			}
		}
	}

	if !haveLetterToInsertAt(firstIdxPointer, sortedLetterInfos) && !haveLetterToInsertAt(secondIdxPointer, sortedLetterInfos) {
		return string(currentResult)
	}

	if haveLetterToInsertAt(firstIdxPointer, sortedLetterInfos) {
		return generateResultByInsertLeftOverFrom(firstIdxPointer, currentResult, sortedLetterInfos)
	} else {
		return generateResultByInsertLeftOverFrom(secondIdxPointer, currentResult, sortedLetterInfos)
	}
}

func categorizeNumberOfEachRuneIn(runes []rune) (letterMapper map[rune]int) {

	letterMapper = make(map[rune]int)

	for _, rune := range runes {

		if _, isExist := letterMapper[rune]; isExist {
			letterMapper[rune]++
		} else {
			letterMapper[rune] = 1
		}
	}

	return letterMapper
}

func createLetterInStringInfosFrom(runeMap map[rune]int) []models.LetterInfo {
	infos := make([]models.LetterInfo, 0)

	for letter, numberOfLetter := range runeMap {
		infos = append(infos, models.LetterInfo{
			Letter:         letter,
			NumberOfLetter: numberOfLetter,
		})
	}

	return infos
}

func descSort(letterInfos []models.LetterInfo) []models.LetterInfo {
	sort.Slice(letterInfos, func(i, j int) bool {
		return letterInfos[i].NumberOfLetter > letterInfos[j].NumberOfLetter
	})

	return letterInfos
}

func haveLetterToInsertAt(idxPointer int, letterInfos []models.LetterInfo) bool {
	return letterInfos[idxPointer].NumberOfLetter > 0
}

func trySelectNextPointerFor(indexToSelect int, letterInfos []models.LetterInfo, ownedIndex int) (newSelectedIndex int, canSelectNewIndex bool) {
	for !haveLetterToInsertAt(indexToSelect, letterInfos) || indexToSelect == ownedIndex {
		indexToSelect++

		if indexToSelect >= len(letterInfos) {
			return -1, false
		}
	}

	return indexToSelect, true
}

func generateResultByInsertLeftOverFrom(pointerToLeftOverLetterInfo int, currentResult []rune, letterInfos []models.LetterInfo) string {
	if result, canInsert := tryInsertLeftOverLettersTo(currentResult, letterInfos[pointerToLeftOverLetterInfo]); canInsert {
		return string(result)
	} else {
		return notPossibleCase
	}
}

func tryInsertLeftOverLettersTo(runes []rune, letterInfo models.LetterInfo) (result []rune, canInsert bool) {
	idxToInsert := 0

	for letterInfo.NumberOfLetter > 0 && idxToInsert <= len(runes) {

		switch idxToInsert {
		case 0:
			if runes[idxToInsert] != letterInfo.Letter {
				runes = append([]rune{letterInfo.Letter}, runes...)

				letterInfo.NumberOfLetter--
			}
		case len(runes):
			if runes[idxToInsert-1] != letterInfo.Letter {
				runes = append(runes, letterInfo.Letter)

				letterInfo.NumberOfLetter--
			}
		default:
			if runes[idxToInsert] != letterInfo.Letter && runes[idxToInsert-1] != letterInfo.Letter {
				runes = append(runes[:idxToInsert+1], runes[idxToInsert:]...)
				runes[idxToInsert] = letterInfo.Letter

				letterInfo.NumberOfLetter--
			}
		}

		idxToInsert = idxToInsert + 2
	}

	if letterInfo.NumberOfLetter > 0 {
		return nil, false
	}

	return runes, true
}
