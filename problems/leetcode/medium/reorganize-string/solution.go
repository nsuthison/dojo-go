package solution

import "sort"

const notPossibleCase string = ""

// Question: https://leetcode.com/problems/reorganize-string/
func reorganizeString(S string) string {

	runes := []rune(S)

	//letterMapper := categorizeNumberOfEachRuneIn(runes)
	//letterInfoes := createLetterInStringInfosFrom(letterMapper)
	//sortedLetterInfoes := descSort(letterInfoes)
	//
	//firstIdxPointer := 0
	//secondIdxPointer := 1
	//
	//for firstIdxPointer < len(sortedLetterInfoes)

	return string(runes)
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

func descSort(letterInfos []lettersInStringInfo) []lettersInStringInfo {
	sort.Slice(letterInfos, func(i,j int) bool {
		return letterInfos[i].numberOfLetter > letterInfos[j].numberOfLetter
	})

	return letterInfos
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

type lettersInStringInfo struct {
	letter rune
	numberOfLetter int
}
