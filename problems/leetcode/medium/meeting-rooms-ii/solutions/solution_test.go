package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var minMeetingRoomsSuccessCaseTestCases = []struct {
	input    [][]int
	expected int
}{
	{[][]int{{0, 30}, {5,10}, {15,20}}, 2},
	{[][]int{{13, 15}, {1,13}}, 1},
}

func Test_minMeetingRooms(t *testing.T) {
	for _, testCase := range minMeetingRoomsSuccessCaseTestCases {
		t.Run("minMeetingRoomsSuccessCase", func(t *testing.T) {
			Convey("Given some time intervals to book meeting room", t, func() {
				input := testCase.input

				Convey("When try to get minimum meeting room to served all booking time", func() {
					result := minMeetingRooms(input)

					Convey("Then minimum number of meeting rooms require should natch expected", func() {
						So(result, ShouldEqual, testCase.expected)
					})
				})
			})
		})
	}
}
