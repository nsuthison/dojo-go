package solutions

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var addStringTestCases = []struct {
	num1     string
	num2     string
	expected string
}{
	{"2", "2", "4"},
	{"5", "212", "217"},
	{"98", "23", "121"},
	{"3876620623801494171", "6529364523802684779", "10405985147604178950"},
}

func Test_addString(t *testing.T) {
	for _, testCase := range addStringTestCases {
		t.Run("addStringTestCases", func(t *testing.T) {
			Convey("Given 2 strings contain only number letter", t, func() {
				num1 := testCase.num1
				num2 := testCase.num2

				Convey("When covert it to int", func() {
					result := addStrings(num1, num2)

					Convey("Then the result should be the summation of these 2 string", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}
