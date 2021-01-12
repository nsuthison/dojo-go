package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var isLeftWordLesserThanRightTestCases = []struct {
	leftWord       string
	rightWord      string
	orderRule      map[rune]int
	expectedResult bool
}{
	{"hello", "leetcode", map[rune]int{'h': 0, 'l': 1}, true},
	{"word", "world", map[rune]int{'w': 0, 'o': 1, 'r': 2, 'l': 3, 'd': 4}, false},
	{"apple", "app", map[rune]int{'a': 0, 'p': 1, 'l': 2, 'e': 3}, false},
}

func Test_createLetterOrderMappingFrom(t *testing.T) {
	Convey("Given string contain letters which represent in order", t, func() {

		input := "bcgha"
		Convey("When create letter order mapping", func() {

			mapper := createLetterOrderMappingFrom(input)
			Convey("Then result of the mapping should reflect given letter in order", func() {
				So(mapper['b'], ShouldEqual, 0)
				So(mapper['c'], ShouldEqual, 1)
				So(mapper['g'], ShouldEqual, 2)
				So(mapper['h'], ShouldEqual, 3)
				So(mapper['a'], ShouldEqual, 4)
			})
		})
	})
}

func Test_isLeftWordLesserThanRight(t *testing.T) {
	for _, testCaseData := range isLeftWordLesserThanRightTestCases {
		t.Run("reorganizeStringSuccessCase", func(t *testing.T) {
			Convey("Given some random string", t, func() {
				leftWord := testCaseData.leftWord
				rightWord := testCaseData.rightWord
				letterOrderRule := testCaseData.orderRule

				Convey("When pass the random string to reorganizeString function", func() {
					isLeftWordLesserThanRight := isLeftWordLesserThanRight(leftWord, rightWord, letterOrderRule)

					Convey("Then the result string shouldn't have same Letter adjacent to each other", func() {
						So(isLeftWordLesserThanRight, ShouldEqual, testCaseData.expectedResult)
					})
				})
			})
		})
	}
}
