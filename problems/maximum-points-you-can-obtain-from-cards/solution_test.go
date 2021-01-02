package solution

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_createSumMapFromFront(t *testing.T) {
	Convey("Given set of cards point and number of card that can be taken", t, func() {
		cards := []int{1,79,80,1,1,1,200,1}
		numberOfCardsToSelect := 3

		Convey("Create mapping summation of cards value starting from front", func() {
			result := createSumMapFromFront(cards, numberOfCardsToSelect)

			Convey("Then it should return array of summation of values base from each index itself plus every element before it", func() {
				So(result[0], ShouldEqual, 1)
				So(result[1], ShouldEqual, 80)
				So(result[2], ShouldEqual, 160)
			})
		})
	})
}