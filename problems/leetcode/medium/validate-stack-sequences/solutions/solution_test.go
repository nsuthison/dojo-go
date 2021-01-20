package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var validateStackSequencesTestCases = []struct {
	pushed         []int
	popped         []int
	expectedResult bool
}{
	{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
	{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	{
		[]int{45, 46, 6, 43, 12, 35, 14, 19, 29, 22, 30, 15, 2, 49, 31, 28, 44, 42, 10, 48, 39, 7, 24, 1, 8, 27, 13, 21, 0, 17, 36, 16, 23, 18, 5, 11, 41, 9, 32, 20, 37, 33, 3, 40, 38, 4, 47, 25, 34, 26},
		[]int{6, 46, 12, 43, 45, 14, 35, 19, 30, 15, 2, 22, 28, 31, 49, 44, 10, 48, 42, 39, 7, 29, 1, 24, 8, 13, 27, 0, 17, 21, 18, 23, 16, 5, 36, 11, 9, 41, 32, 20, 37, 3, 33, 4, 26, 34, 25, 47, 38, 40},
		true,
	},
}

func Test_validateStackSequences(t *testing.T) {
	for _, testCase := range validateStackSequencesTestCases {
		t.Run("validateStackSequences", func(t *testing.T) {
			Convey("Given push and pop", t, func() {

				pushed := testCase.pushed
				popped := testCase.popped

				Convey("When validate stack sequence", func() {

					isStackSequenceValid := validateStackSequences(pushed, popped)

					Convey("Then the result should reflect whether pushed and popped are valid", func() {
						So(isStackSequenceValid, ShouldEqual, testCase.expectedResult)
					})
				})
			})
		})
	}
}
