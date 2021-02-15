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
	//{"3876620623801494171", "6529364523802684779", "10405985147604178950"},
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

var getValueFromTestCases = []struct {
	number      rune
	powerFactor float64
	expected    int
}{
	{'1', 2.0, 100},
	{'9', 0.0, 9},
}

func Test_getValueFrom(t *testing.T) {
	for _, testCase := range getValueFromTestCases {
		t.Run("addStringTestCases", func(t *testing.T) {
			Convey("Given number letter and powerFactor base 10", t, func() {
				number := testCase.number
				powerFactor := testCase.powerFactor

				Convey("When get value", func() {
					result := getValueFrom(number, powerFactor)

					Convey("Then result should be value of number with power factor on base 10", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}

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

					Convey("Then the result should be an integer represent by the given string", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}
