package solution

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

//func Test_maxScore(t *testing.T) {
//	Convey("Given set of cards point and number of card that can be taken", t, func() {
//		cards := []int{1,79,80,1,1,1,200,1}
//		numberOfCardsToSelect := 3
//
//		Convey("When find maxScore", func() {
//			result := maxScore(cards, numberOfCardsToSelect)
//
//			Convey("Then it should return maxScore", func() {
//				So(result, ShouldEqual, 202)
//			})
//		})
//	})
//}

func Test_getSummationOfCardsFromFront_Should_ReturnSummationValueToSelectedCardsAlsoUpdateMemoize_When_MemoizeIsAlreadyHaveValue(t *testing.T) {
	Convey("Given set of cards point and number of card that can be taken", t, func() {
		cards := []int{1,79,80,1,1,1,200,1}
		numberOfCardsToSelect := 5

		Convey("Given memoize", func(){

			memoize := []int{1,80,160}

			Convey("When create mapping summation of cards value starting from front", func() {
				result := getSummationOfCardsFrom(Front, cards, numberOfCardsToSelect, &memoize)

				Convey("Then it should return summation of cards score", func() {
					So(result, ShouldEqual, 162)

					Convey("And memoize should update correctly", func(){
						So(len(memoize), ShouldEqual, 5)
						So(memoize[3], ShouldEqual, 161)
						So(memoize[4], ShouldEqual, 162)
					})
				})
			})
		})
	})
}

func Test_getSummationOfCardsFromFront_Should_ReturnSummationValueToSelectedCardsAlsoUpdateMemoize_When_MemoizeIsEmpty(t *testing.T) {
	Convey("Given set of cards point and number of card that can be taken", t, func() {
		cards := []int{1,79,80,1,1,1,200,1}
		numberOfCardsToSelect := 5

		Convey("Given memoize", func(){

			memoize := make([]int, 0)

			Convey("When create mapping summation of cards value starting from front", func() {
				result := getSummationOfCardsFrom(Front, cards, numberOfCardsToSelect, &memoize)

				Convey("Then it should return summation of cards score", func() {
					So(result, ShouldEqual, 162)

					Convey("And memoize should update correctly", func(){
						So(len(memoize), ShouldEqual, 5)

						So(memoize[0], ShouldEqual, 1)
						So(memoize[1], ShouldEqual, 80)
						So(memoize[2], ShouldEqual, 160)
						So(memoize[3], ShouldEqual, 161)
						So(memoize[4], ShouldEqual, 162)
					})
				})
			})
		})
	})
}

func Test_getSummationOfCardsFromBack_Should_ReturnSummationValueToSelectedCardsAlsoUpdateMemoize_When_MemoizeIsAlreadyHaveValue(t *testing.T) {
	Convey("Given set of cards point and number of card that can be taken", t, func() {
		cards := []int{1,79,80,1,1,1,200,1}
		numberOfCardsToSelect := 5

		Convey("Given memoize", func(){

			memoize := []int{1,200,202}

			Convey("When create mapping summation of cards value starting from front", func() {
				result := getSummationOfCardsFrom(Back, cards, numberOfCardsToSelect, &memoize)

				Convey("Then it should return summation of cards score", func() {
					So(result, ShouldEqual, 204)

					Convey("And memoize should update correctly", func(){
						So(len(memoize), ShouldEqual, 5)
						So(memoize[3], ShouldEqual, 203)
						So(memoize[4], ShouldEqual, 204)
					})
				})
			})
		})
	})
}

func Test_getSummationOfCardsFromBack_Should_ReturnSummationValueToSelectedCardsAlsoUpdateMemoize_When_MemoizeIsEmpty(t *testing.T) {
	Convey("Given set of cards point and number of card that can be taken", t, func() {
		cards := []int{1,79,80,1,1,1,200,1}
		numberOfCardsToSelect := 5

		Convey("Given memoize", func(){

			memoize := make([]int, 0)

			Convey("When create mapping summation of cards value starting from front", func() {
				result := getSummationOfCardsFrom(Back, cards, numberOfCardsToSelect, &memoize)

				Convey("Then it should return summation of cards score", func() {
					So(result, ShouldEqual, 204)

					Convey("And memoize should update correctly", func(){
						So(len(memoize), ShouldEqual, 5)

						So(memoize[0], ShouldEqual, 1)
						So(memoize[1], ShouldEqual, 201)
						So(memoize[2], ShouldEqual, 202)
						So(memoize[3], ShouldEqual, 203)
						So(memoize[4], ShouldEqual, 204)
					})
				})
			})
		})
	})
}

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