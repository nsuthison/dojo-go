package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var backspaceCompareTestCases = []struct {
	first          string
	second         string
	expectedResult bool
}{
	{"ab#c", "ad#c", true},
	{"ab##", "c#d#", true},
	{"a##c", "#a#c", true},
	{"a#c", "b", false},
}

func Test_numDecodings(t *testing.T) {
	for _, testCase := range backspaceCompareTestCases {
		t.Run("backspaceCompareTestCases", func(t *testing.T) {
			Convey("Given two strings", t, func() {
				first := testCase.first
				second := testCase.second

				Convey("When compare both strings", func() {
					result := backspaceCompare(first, second)

					Convey("Then the result should reoresent if both string are the same after compute backspace rule", func() {
						So(result, ShouldEqual, testCase.expectedResult)
					})
				})
			})
		})
	}
}
