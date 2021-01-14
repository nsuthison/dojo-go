package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

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
		t.Run("reorganizeStringSuccessCase", func(t *testing.T) {
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
