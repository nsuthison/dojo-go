package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var toIntFromTestCases = []struct {
	input    string
	expected int
}{
	{"123", 123},
	{"2", 2},
	{"98754", 98754},
}

func Test_toIntFrom(t *testing.T) {
	for _, testCase := range toIntFromTestCases {
		t.Run("toIntFromTestCases", func(t *testing.T) {
			Convey("Given strings contain only number letter", t, func() {
				number := testCase.input

				Convey("When covert it to int", func() {
					result := toIntFrom(number)

					Convey("Then the result should reoresent if both string are the same after compute backspace rule", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}
