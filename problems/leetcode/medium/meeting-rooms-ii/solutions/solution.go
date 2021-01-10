package solutions

import (
	. "github.com/nsuthison/dojo-go/problems/leetcode/medium/meeting-rooms-ii/models"
	"sort"
)

// Question: https://leetcode.com/problems/meeting-rooms-ii/
func minMeetingRooms(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	meetingRooms := make([]MeetingRoom, 0)

	for _, interval := range intervals {
		toBook := TimeInterval{
			Start: interval[0],
			End:   interval[1],
		}

		isAbleToBook := false
		meetingRoomIdx := 0

		for meetingRoomIdx < len(meetingRooms) && !isAbleToBook {

			if meetingRoomIdx == len(meetingRooms) {
				meetingRooms = append(meetingRooms, MeetingRoom{
					BookedIntervals: []TimeInterval{toBook},
				})

				break
			}

			meetingRooms[meetingRoomIdx], isAbleToBook = tryBook(toBook, meetingRooms[meetingRoomIdx])

			meetingRoomIdx++
		}

		if !isAbleToBook {
			meetingRooms = append(meetingRooms, MeetingRoom{
				BookedIntervals: []TimeInterval{toBook},
			})
		}
	}

	return len(meetingRooms)
}

func tryBook(toBook TimeInterval, meetingRoom MeetingRoom) (result MeetingRoom, canBook bool) {

	booked := &meetingRoom.BookedIntervals
	for idx := 0; idx <= len(*booked); idx++ {

		switch idx {
		case 0:
			if toBook.IsBefore((*booked)[idx]) {
				*booked = append([]TimeInterval{toBook}, *booked...)

				return meetingRoom, true
			}
		case len(*booked):
			if toBook.IsAfter((*booked)[idx-1]) {
				*booked = append(*booked, toBook)

				return meetingRoom, true
			}
		default:
			if toBook.IsBetween((*booked)[idx-1], (*booked)[idx]) {
				*booked = append((*booked)[:idx+1], (*booked)[idx:]...)
				(*booked)[idx] = toBook

				return meetingRoom, true
			}
		}
	}

	return meetingRoom, false
}
