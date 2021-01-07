package solution

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reorganizeString(t *testing.T) {
	Convey("Given some random string", t, func() {
		input := "aab"

		Convey("When pass the random string to reorganizeString function", func() {
			result := reorganizeString(input)

			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
				for i := 0; i < len(result)-1; i++ {
					So(result[i], ShouldNotEqual, result[i+1])
				}			})
		})
	})
}

func Test_reorganizeString2(t *testing.T) {
	Convey("Given some random string", t, func() {
		input := "baaba"

		Convey("When pass the random string to reorganizeString function", func() {
			result := reorganizeString(input)

			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
				for i := 0; i < len(result)-1; i++ {
					So(result[i], ShouldNotEqual, result[i+1])
				}
			})
		})
	})
}

func Test_reorganizeString3(t *testing.T) {
	Convey("Given some random string", t, func() {
		input := "cxmwmmm"

		Convey("When pass the random string to reorganizeString function", func() {
			result := reorganizeString(input)

			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
				for i := 0; i < len(result)-1; i++ {
					So(result[i], ShouldNotEqual, result[i+1])
				}
			})
		})
	})
}

func Test_reorganizeString4(t *testing.T) {
	Convey("Given some random string", t, func() {
		input := "bbbbbbb"

		Convey("When pass the random string to reorganizeString function", func() {
			result := reorganizeString(input)

			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
				So(result, ShouldEqual, "")
			})
		})
	})
}

func Test_swapRunes(t *testing.T) {
	Convey("Given runes and 2 index", t, func() {
		someString := "someString"
		runes := []rune(someString)

		firstIndexToSwap := 3
		secondIndexToSwap := 4

		Convey("When swap these 2 runes", func() {
			swapRunes(&runes, firstIndexToSwap, secondIndexToSwap)

			Convey("Then runes in both index position should be swap", func() {
				So(string(runes), ShouldEqual, "somSetring")
			})
		})
	})
}

func Test_categorizeNumberOfEachRuneIn(t *testing.T) {
	Convey("Given string", t, func() {
		someString := "cxmwmmm"
		runes := []rune(someString)

		Convey("When categorize number of each rune", func() {
			categorize := categorizeNumberOfEachRuneIn(runes)

			Convey("Then the categorize should ", func() {
				So(categorize['c'], ShouldEqual, 1)
				So(categorize['x'], ShouldEqual, 1)
				So(categorize['w'], ShouldEqual, 1)
				So(categorize['m'], ShouldEqual, 4)
			})
		})
	})
}

func Test_createLetterInStringInfosFrom(t *testing.T) {
	Convey("Given map[rune]int", t, func() {
		someString := "cxmwwmmmwc"
		runes := []rune(someString)

		categorize := categorizeNumberOfEachRuneIn(runes)

		Convey("When create letterInStringInfos", func() {

			result := createLetterInStringInfosFrom(categorize)

			Convey("Then it should create slice with letterInfo", func() {

				for _, letterInfo := range result {
					switch letterInfo.letter {
					case 'x':
						So(letterInfo.numberOfLetter, ShouldEqual, 1)
					case 'w':
						So(letterInfo.numberOfLetter, ShouldEqual, 3)
					case 'm':
						So(letterInfo.numberOfLetter, ShouldEqual, 4)
					case 'c':
						So(letterInfo.numberOfLetter, ShouldEqual, 2)
					}
				}
			})
		})
	})
}

func Test_descSort(t *testing.T) {
	Convey("Given letterInStringInfos", t, func() {
		someString := "cxmwwmmmwc"
		runes := []rune(someString)

		categorize := categorizeNumberOfEachRuneIn(runes)
		letterInfos := createLetterInStringInfosFrom(categorize)

		Convey("When descSort letterInfos", func() {

			sortedLetterInfos := descSort(letterInfos)

			Convey("Then it should create slice with letterInfo", func() {

				So(sortedLetterInfos[0].letter, ShouldEqual, 'm')
				So(sortedLetterInfos[1].letter, ShouldEqual, 'w')
				So(sortedLetterInfos[2].letter, ShouldEqual, 'c')
				So(sortedLetterInfos[3].letter, ShouldEqual, 'x')
			})
		})
	})
}
