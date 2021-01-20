package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var isAlienSortedTestCases = []struct {
	words           []string
	letterOrderRule string
	expectedResult  bool
}{
	{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
	{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
	{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
}

func Test_isAlienSorted(t *testing.T) {
	for _, testCaseData := range isAlienSortedTestCases {
		t.Run("isAlienSorted", func(t *testing.T) {
			Convey("Given words and letter order rule", t, func() {
				words := testCaseData.words
				letterOrderRule := testCaseData.letterOrderRule

				Convey("When check is word are sorted lexicographicaly by given rule", func() {
					isAlienSorted := isAlienSorted(words, letterOrderRule)

					Convey("Then the result should reflect rule as expected", func() {
						So(isAlienSorted, ShouldEqual, testCaseData.expectedResult)
					})
				})
			})
		})
	}
}

func Test_createLetterOrderMappingFrom(t *testing.T) {
	Convey("Given string contain letters which represent in order", t, func() {

		input := "bcgha"
		Convey("When create letter order mapping", func() {

			mapper := createLetterOrderMappingRuleFrom(input)
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

func Test_isLeftWordLesserThanRight(t *testing.T) {
	for _, testCaseData := range isLeftWordLesserThanRightTestCases {
		t.Run("reorganizeStringSuccessCase", func(t *testing.T) {
			Convey("Given left and right word to compare with letter order rule", t, func() {
				leftWord := testCaseData.leftWord
				rightWord := testCaseData.rightWord
				letterOrderRule := testCaseData.orderRule

				Convey("When checking is left word lesser than right word by rule", func() {
					isLeftWordLesserThanRight := isLeftWordLesserThanRight(leftWord, rightWord, letterOrderRule)

					Convey("Then the result should reflect given rule", func() {
						So(isLeftWordLesserThanRight, ShouldEqual, testCaseData.expectedResult)
					})
				})
			})
		})
	}
}
