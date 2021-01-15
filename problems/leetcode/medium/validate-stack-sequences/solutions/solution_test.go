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

func Test_Push(t *testing.T) {
	Convey("Given stack and element to push", t, func() {
		stack := make(Stack, 0)
		toPush := 3

		Convey("When push element to stack", func() {
			stack.Push(toPush)

			Convey("Then last element of stack should be the element that just push to stack", func() {
				So(stack[len(stack)-1], ShouldEqual, toPush)
			})
		})
	})
}

func Test_Pop_SuccessCase(t *testing.T) {
	Convey("Given stack with element", t, func() {
		stack := Stack{14, 3, 7}

		Convey("When pop element from stack", func() {
			pop, canPop := stack.Pop()

			Convey("Then it should return last element from stack", func() {
				So(canPop, ShouldBeTrue)

				So(pop, ShouldEqual, 7)
				So(len(stack), ShouldEqual, 2)
				So(stack[0], ShouldEqual, 14)
				So(stack[1], ShouldEqual, 3)
			})
		})
	})
}

func Test_Pop_FailCase(t *testing.T) {
	Convey("Given empty stack", t, func() {
		stack := Stack{}

		Convey("When pop element from stack", func() {
			_, canPop := stack.Pop()

			Convey("Then canPop should be fault", func() {
				So(canPop, ShouldBeFalse)
			})
		})
	})
}
