package solutions

import (
	"sort"

	models "github.com/nsuthison/dojo-go/problems/leetcode/medium/meeting-rooms-ii/models"
)

// Question: https://leetcode.com/problems/meeting-rooms-ii/
func minMeetingRooms(intervals [][]int) int {

	ascSort(&intervals)

	meetingRooms := make([]models.MeetingRoom, 0)

	for _, interval := range intervals {

		toBook := models.TimeInterval{
			Start: interval[0],
			End:   interval[1],
		}

		var isAbleToBook bool
		meetingRooms, isAbleToBook = tryToFindRoomAndBookWith(toBook, meetingRooms)

		if !isAbleToBook {
			meetingRooms = createNewMeetingRoomAndBookWith(toBook, meetingRooms)
		}
	}

	return len(meetingRooms)
}

func ascSort(intervals *[][]int) {
	sort.Slice(*intervals, func(i, j int) bool {
		return (*intervals)[i][0] < (*intervals)[j][0]
	})
}

func tryToFindRoomAndBookWith(toBook models.TimeInterval, meetingRooms []models.MeetingRoom) (result []models.MeetingRoom, canBook bool) {
	meetingRoomIdx := 0
	isAbleToBook := false

	for meetingRoomIdx < len(meetingRooms) && !isAbleToBook {
		meetingRooms[meetingRoomIdx], isAbleToBook = tryBook(toBook, meetingRooms[meetingRoomIdx])

		meetingRoomIdx++
	}

	return meetingRooms, isAbleToBook
}

func tryBook(toBook models.TimeInterval, meetingRoom models.MeetingRoom) (result models.MeetingRoom, canBook bool) {

	booked := &meetingRoom.BookedIntervals
	for idx := 0; idx <= len(*booked); idx++ {

		switch idx {
		case 0:
			if toBook.IsBefore((*booked)[idx]) {
				*booked = append([]models.TimeInterval{toBook}, *booked...)

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

func createNewMeetingRoomAndBookWith(toBook models.TimeInterval, meetingRooms []models.MeetingRoom) []models.MeetingRoom {
	return append(meetingRooms, models.MeetingRoom{
		BookedIntervals: []models.TimeInterval{toBook},
	})
}
