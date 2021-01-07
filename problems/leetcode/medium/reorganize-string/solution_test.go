package solution

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//func Test_reorganizeString(t *testing.T) {
//	Convey("Given some random string", t, func() {
//		input := "aab"
//
//		Convey("When pass the random string to reorganizeString function", func() {
//			result := reorganizeString(input)
//
//			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
//				So(result, ShouldEqual, "aba")
//			})
//		})
//	})
//}
//
//func Test_reorganizeString2(t *testing.T) {
//	Convey("Given some random string", t, func() {
//		input := "baaba"
//
//		Convey("When pass the random string to reorganizeString function", func() {
//			result := reorganizeString(input)
//
//			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
//				So(result, ShouldEqual, "ababa")
//			})
//		})
//	})
//}

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

//func Test_createLetterInStringInfosFrom(t *testing.T) {
//	Convey("Given map[rune]int", t, func() {
//		someString := "cxmwwmmwmwc"
//		runes := []rune(someString)
//
//		categorize := categorizeNumberOfEachRuneIn(runes)
//
//		Convey("When create letterInStringInfos", func() {
//
//			result := createLetterInStringInfosFrom(categorize)
//
//			Convey("Then the categorize should ", func() {
//
//				So(result, ShouldEqual, 1)
//				So(categorize['x'], ShouldEqual, 1)
//				So(categorize['w'], ShouldEqual, 1)
//				So(categorize['m'], ShouldEqual, 4)
//			})
//		})
//	})
//}