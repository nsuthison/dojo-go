package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_createLetterOrderMappingFrom(t *testing.T) {
	Convey("Given string contain letters which represent in order",t , func() {

		input := "bcgha"
		Convey("When create letter order mapping", func() {

			mapper := createLetterOrderMappingFrom(input)
			Convey("Then result of the mapping should reflect given letter in order", func() {
				So(mapper['b'], ShouldEqual, 0)
				So(mapper['c'], ShouldEqual, 1)
				So(mapper['g'], ShouldEqual, 2)
				So(mapper['h'], ShouldEqual, 3)
				So(mapper['a'], ShouldEqual, 4)
			})
		})
	})
}
