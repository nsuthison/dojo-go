package solution

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reorganizeString(t *testing.T) {
	Convey("Given some random string", t, func() {
		input := "aab"

		Convey("When pass the random string to reorganizeString function", func() {
			result := reorganizeString(input)

			Convey("Then the result string shouldn't have same letter adjacent to each other", func() {
				So(result, ShouldEqual, "aba")
			})
		})
	})
}