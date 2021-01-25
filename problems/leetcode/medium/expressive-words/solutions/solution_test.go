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
	{"", []string{"hello", "hi", "helo"}, 0},
	{"abcd", []string{"abc"}, 0},
	{"aaa", []string{"aaaa"}, 0},
	{"le", []string{"lee"}, 0},
}

func Test_expressiveWords(t *testing.T) {
	for _, testCase := range expressiveWordsTestCases {
		t.Run("expressiveWordsTestCases", func(t *testing.T) {
			Convey("Given input string and words", t, func() {
				input := testCase.input
				words := testCase.words

				Convey("When try to find how many words are stretchy", func() {
					result := expressiveWords(input, words)

					Convey("Then the result number should return number of words that are recognize as stretchy", func() {
						So(result, ShouldEqual, testCase.expectedResult)
					})
				})
			})
		})
	}
}

func Test_createStretchyInfosFrom(t *testing.T) {
	Convey("Given some string", t, func() {

		someString := "heeellooo"

		Convey("When create stretchyInfos", func() {

			stretchyMap := createStretchyInfosFrom(someString)

			Convey("Then it should create stretchyInfos in order", func() {
				So(stretchyMap[0].Letter, ShouldEqual, 'h')
				So(stretchyMap[0].IsStretchy, ShouldEqual, false)
				So(stretchyMap[0].LetterCount, ShouldEqual, 1)

				So(stretchyMap[1].Letter, ShouldEqual, 'e')
				So(stretchyMap[1].IsStretchy, ShouldEqual, true)
				So(stretchyMap[1].LetterCount, ShouldEqual, 3)

				So(stretchyMap[2].Letter, ShouldEqual, 'l')
				So(stretchyMap[2].IsStretchy, ShouldEqual, false)
				So(stretchyMap[2].LetterCount, ShouldEqual, 1)

				So(stretchyMap[3].Letter, ShouldEqual, 'l')
				So(stretchyMap[3].IsStretchy, ShouldEqual, false)
				So(stretchyMap[3].LetterCount, ShouldEqual, 1)

				So(stretchyMap[4].Letter, ShouldEqual, 'o')
				So(stretchyMap[4].IsStretchy, ShouldEqual, true)
				So(stretchyMap[4].LetterCount, ShouldEqual, 3)
			})
		})
	})
}

func Test_isWordStretchy(t *testing.T) {
	Convey("Given stretchyInfos and string which valid to the infos", t, func() {

		someString := "hello"
		stretchyInfos := createStretchyInfosFrom("heeellooo")

		Convey("When check if the given word is stretchy", func() {

			isWordStretchy := isWordStretchy(someString, stretchyInfos)

			Convey("Then result should be true", func() {
				So(isWordStretchy, ShouldBeTrue)
			})
		})
	})
}
