package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var numDecodingsTestCases = []struct {
	input          string
	expectedResult int
}{
	{"12",2},
	{"226",3},
	{"0", 0},
	{"2101", 1},
	{"10011", 0},
}

func Test_numDecodings(t *testing.T) {
	for _, testCase := range numDecodingsTestCases {
		t.Run("numDecodingsTestCases", func(t *testing.T) {
			Convey("Given input string contain only digits", t, func() {
				input := testCase.input

				Convey("When pass the random string to reorganizeString function", func() {
					result := numDecodings(input)

					Convey("Then the result string shouldn't have same Letter adjacent to each other", func() {
						So(result, ShouldEqual, testCase.expectedResult)
					})
				})
			})
		})
	}
}

var isAbleToCombineTestCases = []struct {
	head           int
	tail           int
	expectedResult bool
}{
	{1, 2, true},
	{2, 6, true},
	{2, 0, true},
	{5, 2, false},
	{0, 6, false},
	{2, 7, false},
}

func Test_isBeAbleToCombine(t *testing.T) {
	for _, testCase := range isAbleToCombineTestCases {
		t.Run("isAbleToCombineTestCases", func(t *testing.T) {
			Convey("Given head and tail as int to combine", t, func() {
				head := testCase.head
				tail := testCase.tail

				Convey("When pass the random string to reorganizeString function", func() {
					result := isBeAbleToCombine(head, tail)

					Convey("Then the result string shouldn't have same Letter adjacent to each other", func() {
						So(result, ShouldEqual, testCase.expectedResult)
					})
				})
			})
		})
	}
}
