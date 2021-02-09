package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var fullJustifyTestCases = []struct {
	words          []string
	maxWidth       int
	expectedResult []string
}{
	{
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
	},
}

func Test_fullJustify(t *testing.T) {
	for _, testCase := range fullJustifyTestCases {
		t.Run("fullJustify test", func(t *testing.T) {
			Convey("Given words and maxwidth to justify", t, func() {
				words := testCase.words
				maxWidth := testCase.maxWidth

				Convey("When justify", func() {
					result := fullJustify(words, maxWidth)

					Convey("Then the result string of each line should be fit with maxWidth", func() {
						for i := 0; i < len(testCase.expectedResult); i++ {
							So(result[i], ShouldEqual, testCase.expectedResult[i])
						}
					})
				})
			})
		})
	}
}

func Test_groupWordsToLineFrom(t *testing.T) {
	Convey("Given words and maximumWidth", t, func() {
		words := []string{"This", "is", "an", "example", "of", "text", "justification."}
		maxWidth := 16

		Convey("When group words to line", func() {
			result := groupWordsToLineFrom(words, maxWidth)

			Convey("Then the result should be words separate to each line", func() {
				expected := [][]string{
					{"This", "is", "an"},
					{"example", "of", "text"},
					{"justification."},
				}

				for lineIdx := 0; lineIdx < len(result); lineIdx++ {
					for wordIdx := 0; wordIdx < len(result[lineIdx]); wordIdx++ {
						So(result[lineIdx][wordIdx], ShouldEqual, expected[lineIdx][wordIdx])
					}
				}
			})
		})
	})
}

var createJustifyTextTestCases = []struct {
	words                 []string
	spaceWidthBetweenWord int
	spaceFraction         int
	expected              string
}{
	{
		[]string{"This", "is", "an"},
		4,
		0,
		"This    is    an",
	},
	{
		[]string{"example", "of", "text"},
		1,
		1,
		"example  of text",
	},
	{
		[]string{"justification."},
		2,
		0,
		"justification.  ",
	},
}

func Test_createJustifyTextFrom(t *testing.T) {

	for _, testCase := range createJustifyTextTestCases {
		t.Run("Test_createJustifyTextFrom test", func(t *testing.T) {
			Convey("Given words, number of space between each word, and number of space fraction", t, func() {
				words := testCase.words
				spaceWidthBetweenWord := testCase.spaceWidthBetweenWord
				spaceFraction := testCase.spaceFraction

				Convey("When create justify text", func() {
					result := createJustifyTextFrom(words, spaceWidthBetweenWord, spaceFraction)

					Convey("Then the result should be the text that combine each word with spaces", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}
