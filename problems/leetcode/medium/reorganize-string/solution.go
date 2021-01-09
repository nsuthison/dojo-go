package solution

import (
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
	letterInfoes := createLetterInStringInfosFrom(letterMapper)
	sortedLetterInfoes := descSort(letterInfoes)

	firstIdxPointer := 0
	secondIdxPointer := 1

	result := make([]rune, 0)

	if len(sortedLetterInfoes) <= 1 {
		return notPossibleCase
	}
	
	for firstIdxPointer < len(sortedLetterInfoes) {
		for sortedLetterInfoes[firstIdxPointer].numberOfLetter > 0 && sortedLetterInfoes[secondIdxPointer].numberOfLetter > 0 {
			result = append(result, sortedLetterInfoes[firstIdxPointer].letter)
			result = append(result, sortedLetterInfoes[secondIdxPointer].letter)

			sortedLetterInfoes[firstIdxPointer].numberOfLetter--
			sortedLetterInfoes[secondIdxPointer].numberOfLetter--
		}

		if sortedLetterInfoes[firstIdxPointer].numberOfLetter <= 0 {
			if newSelectedPointer, canSelectNewIndex := selectNextPointer(sortedLetterInfoes, firstIdxPointer, secondIdxPointer); canSelectNewIndex {
				firstIdxPointer = newSelectedPointer
			} else {
				if sortedLetterInfoes[secondIdxPointer].numberOfLetter == 0 {
					return string(result)
				} else {
					if result, canInsert := tryInsertLeftOverLettersTo(result, sortedLetterInfoes[secondIdxPointer]); canInsert {
						return string(result)
					} else {
						return notPossibleCase
					}
				}
			}
		}

		if sortedLetterInfoes[secondIdxPointer].numberOfLetter <= 0 {
			if newSelectedPointer, canSelectNewIndex := selectNextPointer(sortedLetterInfoes, secondIdxPointer, firstIdxPointer); canSelectNewIndex {
				secondIdxPointer = newSelectedPointer
			} else {
				if sortedLetterInfoes[firstIdxPointer].numberOfLetter == 0 {
					return string(result)
				} else {
					if result, canInsert := tryInsertLeftOverLettersTo(result, sortedLetterInfoes[firstIdxPointer]); canInsert {
						return string(result)
					} else {
						return notPossibleCase
					}
				}
			}
		}
	}

	return string(runes)
}

func selectNextPointer(letterInfos []lettersInStringInfo, indexToSelect int, anotherIdx int) (newSelectedIndex int, canSelectNewIndex bool) {
	for letterInfos[indexToSelect].numberOfLetter == 0 || indexToSelect == anotherIdx {
		indexToSelect++

		if indexToSelect >= len(letterInfos) {
			return -1, false
		}
	}

	return indexToSelect, true
}

func swapRunes(runes *[]rune, firstIndex int, secondIndex int) {
	(*runes)[firstIndex], (*runes)[secondIndex] = (*runes)[secondIndex], (*runes)[firstIndex]
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

func createLetterInStringInfosFrom(runeMap map[rune]int) []lettersInStringInfo {
	infos := make([]lettersInStringInfo, 0)

	for letter, numberOfLetter := range runeMap {
		infos = append(infos, lettersInStringInfo{
			letter: letter,
			numberOfLetter: numberOfLetter,
		})
	}

	return infos
}

func descSort(letterInfos []lettersInStringInfo) []lettersInStringInfo {
	sort.Slice(letterInfos, func(i,j int) bool {
		return letterInfos[i].numberOfLetter > letterInfos[j].numberOfLetter
	})

	return letterInfos
}

func tryInsertLeftOverLettersTo(runes []rune, letterInfo lettersInStringInfo) (result []rune, canInsert bool) {
	idxToInsert := 0

	for letterInfo.numberOfLetter > 0 && idxToInsert <= len(runes)  {
		if idxToInsert == len(runes) {
			runes = append(runes, letterInfo.letter)
			letterInfo.numberOfLetter--
			break
		}

		if runes[idxToInsert] != letterInfo.letter {
			if idxToInsert > 0 {
				if runes[idxToInsert - 1] != letterInfo.letter {
					runes = append(runes[:idxToInsert + 1], runes[idxToInsert:]...)
					runes[idxToInsert] = letterInfo.letter

					letterInfo.numberOfLetter--
				}
			} else {
				runes = append([]rune{letterInfo.letter}, runes...)

				letterInfo.numberOfLetter--
			}
		}

		idxToInsert = idxToInsert + 2
	}

	if letterInfo.numberOfLetter > 0 {
		return nil, false
	}

	return runes, true
}

type lettersInStringInfo struct {
	letter rune
	numberOfLetter int
}
