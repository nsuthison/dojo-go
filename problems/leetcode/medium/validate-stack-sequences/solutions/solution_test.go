package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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

func Test_Pop(t *testing.T) {
	Convey("Given stack with element", t, func() {
		stack := Stack{14, 7}

		Convey("When push element to stack", func() {
			pop, canPop := stack.Pop()

			Convey("Then last element of stack should be the element that just push to stack", func() {
				So(canPop, ShouldEqual, true)
				So(pop, ShouldEqual, 7)
			})
		})
	})
}
