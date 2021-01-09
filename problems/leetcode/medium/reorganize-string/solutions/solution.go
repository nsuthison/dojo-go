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
	letterInfoes := createLetterInStringInfosFrom(letterMapper)
	sortedLetterInfoes := descSort(letterInfoes)

	firstIdxPointer := 0
	secondIdxPointer := 1

	result := make([]rune, 0)

	if len(sortedLetterInfoes) <= 1 {
		return notPossibleCase
	}
	
	for firstIdxPointer < len(sortedLetterInfoes) {
		for sortedLetterInfoes[firstIdxPointer].NumberOfLetter > 0 && sortedLetterInfoes[secondIdxPointer].NumberOfLetter > 0 {
			result = append(result, sortedLetterInfoes[firstIdxPointer].Letter)
			result = append(result, sortedLetterInfoes[secondIdxPointer].Letter)

			sortedLetterInfoes[firstIdxPointer].NumberOfLetter--
			sortedLetterInfoes[secondIdxPointer].NumberOfLetter--
		}

		if sortedLetterInfoes[firstIdxPointer].NumberOfLetter <= 0 {
			if newSelectedPointer, canSelectNewIndex := selectNextPointer(sortedLetterInfoes, firstIdxPointer, secondIdxPointer); canSelectNewIndex {
				firstIdxPointer = newSelectedPointer
			} else {
				if sortedLetterInfoes[secondIdxPointer].NumberOfLetter == 0 {
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

		if sortedLetterInfoes[secondIdxPointer].NumberOfLetter <= 0 {
			if newSelectedPointer, canSelectNewIndex := selectNextPointer(sortedLetterInfoes, secondIdxPointer, firstIdxPointer); canSelectNewIndex {
				secondIdxPointer = newSelectedPointer
			} else {
				if sortedLetterInfoes[firstIdxPointer].NumberOfLetter == 0 {
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

func selectNextPointer(letterInfos []models.LetterInfo, indexToSelect int, anotherIdx int) (newSelectedIndex int, canSelectNewIndex bool) {
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

func tryInsertLeftOverLettersTo(runes []rune, letterInfo models.LetterInfo) (result []rune, canInsert bool) {
	idxToInsert := 0

	for letterInfo.NumberOfLetter > 0 && idxToInsert <= len(runes)  {
		if idxToInsert == len(runes) {
			runes = append(runes, letterInfo.Letter)
			letterInfo.NumberOfLetter--
			break
		}

		if runes[idxToInsert] != letterInfo.Letter {
			if idxToInsert > 0 {
				if runes[idxToInsert - 1] != letterInfo.Letter {
					runes = append(runes[:idxToInsert + 1], runes[idxToInsert:]...)
					runes[idxToInsert] = letterInfo.Letter

					letterInfo.NumberOfLetter--
				}
			} else {
				runes = append([]rune{letterInfo.Letter}, runes...)

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