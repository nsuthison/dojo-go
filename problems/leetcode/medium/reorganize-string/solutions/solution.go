package solutions

import (
	"github.com/nsuthison/dojo-go/problems/leetcode/medium/reorganize-string/models"
	"sort"
)

const notPossibleCase string = ""

// Question: https://leetcode.com/problems/reorganize-string/

/// Solution is that, sort it first.
func reorganizeString(S string) string {

	runes := []rune(S)

	if len(S) <= 1 {
		return S
	}

	letterMapper := categorizeNumberOfEachRuneIn(runes)
	letterInfos := createLetterInStringInfosFrom(letterMapper)
	sortedLetterInfos := descSort(letterInfos)

	firstIdxPointer := 0
	secondIdxPointer := 1

	result := make([]rune, 0)

	if len(sortedLetterInfos) <= 1 {
		return notPossibleCase
	}
	
	for firstIdxPointer < len(sortedLetterInfos) {
		for haveLetterToInsertAt(firstIdxPointer, sortedLetterInfos) && haveLetterToInsertAt(secondIdxPointer, sortedLetterInfos) {

			result = append(result, sortedLetterInfos[firstIdxPointer].Letter)
			result = append(result, sortedLetterInfos[secondIdxPointer].Letter)

			sortedLetterInfos[firstIdxPointer].NumberOfLetter--
			sortedLetterInfos[secondIdxPointer].NumberOfLetter--
		}

		if sortedLetterInfos[firstIdxPointer].NumberOfLetter <= 0 {
			newSelectedPointer, canSelectNextPointer := trySelectNextPointer(sortedLetterInfos, firstIdxPointer, secondIdxPointer)

			if canSelectNextPointer {
				firstIdxPointer = newSelectedPointer
			} else {
				return generateResult(result, sortedLetterInfos, secondIdxPointer)
			}
		}

		if sortedLetterInfos[secondIdxPointer].NumberOfLetter <= 0 {
			newSelectedPointer, canSelectNextPointer := trySelectNextPointer(sortedLetterInfos, secondIdxPointer, firstIdxPointer)

			if canSelectNextPointer {
				secondIdxPointer = newSelectedPointer
			} else {
				return generateResult(result, sortedLetterInfos, firstIdxPointer)
			}
		}
	}

	return string(runes)
}

func haveLetterToInsertAt(idxPointer int, letterInfos []models.LetterInfo) bool {
	return letterInfos[idxPointer].NumberOfLetter > 0
}

func trySelectNextPointer(letterInfos []models.LetterInfo, indexToSelect int, anotherIdx int) (newSelectedIndex int, canSelectNewIndex bool) {
	for letterInfos[indexToSelect].NumberOfLetter == 0 || indexToSelect == anotherIdx {
		indexToSelect++

		if indexToSelect >= len(letterInfos) {
			return -1, false
		}
	}

	return indexToSelect, true
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
	sort.Slice(letterInfos, func(i,j int) bool {
		return letterInfos[i].NumberOfLetter > letterInfos[j].NumberOfLetter
	})

	return letterInfos
}

func generateResult(currentResult []rune, letterInfos []models.LetterInfo, pointerToLeftOverLetterInfo int) string {

	if letterInfos[pointerToLeftOverLetterInfo].NumberOfLetter == 0 {
		return string(currentResult)
	} else {
		if result, canInsert := tryInsertLeftOverLettersTo(currentResult, letterInfos[pointerToLeftOverLetterInfo]); canInsert {
			return string(result)
		} else {
			return notPossibleCase
		}
	}
}

func tryInsertLeftOverLettersTo(runes []rune, letterInfo models.LetterInfo) (result []rune, canInsert bool) {
	idxToInsert := 0

	for letterInfo.NumberOfLetter > 0 && idxToInsert <= len(runes)  {

		switch idxToInsert {
		case 0:
			if runes[idxToInsert] != letterInfo.Letter {
				runes = append([]rune{letterInfo.Letter}, runes...)

				letterInfo.NumberOfLetter--
			}
		case len(runes):
			runes = append(runes, letterInfo.Letter)
			letterInfo.NumberOfLetter--
		default:
			if runes[idxToInsert] != letterInfo.Letter && runes[idxToInsert - 1] != letterInfo.Letter {
				runes = append(runes[:idxToInsert + 1], runes[idxToInsert:]...)
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