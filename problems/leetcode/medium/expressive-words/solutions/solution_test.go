package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var expressiveWordsTestCases = []struct {
	input          string
	words          []string
	expectedResult int
}{
	{"heeellooo", []string{"hello", "hi", "helo"}, 1},
}

// func Test_expressiveWords(t *testing.T) {
// 	for _, testCase := range expressiveWordsTestCases {
// 		t.Run("expressiveWordsTestCases", func(t *testing.T) {
// 			Convey("Given input string and words", t, func() {
// 				input := testCase.input
// 				words := testCase.words

// 				Convey("When try to find how many words are stretchy", func() {
// 					result := expressiveWords(input, words)

// 					Convey("Then the result number should return number of words that are recognize as stretchy", func() {
// 						So(result, ShouldEqual, testCase.expectedResult)
// 					})
// 				})
// 			})
// 		})
// 	}
// }

func Test_createStretchyMap(t *testing.T) {
	Convey("Given some string", t, func() {

		someString := "heeellooo"

		Convey("", func() {

			stretchyMap := createStretchyMap(someString)

			Convey("", func() {
				So(stretchyMap[0].Letter, ShouldEqual, 'h')
				So(stretchyMap[0].IsStretchy, ShouldEqual, false)

				So(stretchyMap[1].Letter, ShouldEqual, 'e')
				So(stretchyMap[1].IsStretchy, ShouldEqual, true)

				So(stretchyMap[2].Letter, ShouldEqual, 'l')
				So(stretchyMap[2].IsStretchy, ShouldEqual, false)

				So(stretchyMap[3].Letter, ShouldEqual, 'l')
				So(stretchyMap[3].IsStretchy, ShouldEqual, false)

				So(stretchyMap[4].Letter, ShouldEqual, 'o')
				So(stretchyMap[4].IsStretchy, ShouldEqual, true)
			})
		})
	})
}
