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