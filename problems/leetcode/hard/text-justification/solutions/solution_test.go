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
	{
		[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20,
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
	},
	{
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	},
	{
		[]string{"Listen", "to", "many,", "speak", "to", "a", "few."},
		6,
		[]string{
			"Listen",
			"to    ",
			"many, ",
			"speak ",
			"to   a",
			"few.  ",
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

var groupWordsToLineFromTestCases = []struct {
	words    []string
	maxWidth int
	expected [][]string
}{
	{
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[][]string{
			{"This", "is", "an"},
			{"example", "of", "text"},
			{"justification."},
		},
	},
	{
		[]string{"Listen", "to", "many,", "speak", "to", "a", "few."},
		6,
		[][]string{
			{"Listen"},
			{"to"},
			{"many,"},
			{"speak"},
			{"to", "a"},
			{"few."},
		},
	},
}

func Test_groupWordsToLineFrom(t *testing.T) {
	for _, testCase := range groupWordsToLineFromTestCases {
		t.Run("Test_createJustifyTextFrom test", func(t *testing.T) {
			Convey("Given words and maximumWidth", t, func() {
				words := testCase.words
				maxWidth := testCase.maxWidth

				Convey("When group words to line", func() {
					result := groupWordsToLineFrom(words, maxWidth)

					Convey("Then the result should be words separate to each line", func() {

						for lineIdx := 0; lineIdx < len(result); lineIdx++ {
							for wordIdx := 0; wordIdx < len(result[lineIdx]); wordIdx++ {
								So(result[lineIdx][wordIdx], ShouldEqual, testCase.expected[lineIdx][wordIdx])
							}
						}
					})
				})
			})
		})
	}
}

var createJustifyTextTestCases = []struct {
	words    []string
	maxWidth int
	expected string
}{
	{
		[]string{"This", "is", "an"},
		16,
		"This    is    an",
	},
	{
		[]string{"example", "of", "text"},
		16,
		"example  of text",
	},
	{
		[]string{"justification."},
		16,
		"justification.  ",
	},
	{
		[]string{"Science", "is", "what", "we"},
		20,
		"Science  is  what we",
	},
	{
		[]string{"Listen"},
		6,
		"Listen",
	},
}

func Test_createJustifyTextFrom(t *testing.T) {

	for _, testCase := range createJustifyTextTestCases {
		t.Run("Test_createJustifyTextFrom test", func(t *testing.T) {
			Convey("Given words and maxWidth to justify", t, func() {
				words := testCase.words
				maxWidth := testCase.maxWidth

				Convey("When create justify text", func() {
					result := createJustifyTextFrom(words, maxWidth)

					Convey("Then the result should be the text that combine each word with spaces", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}
