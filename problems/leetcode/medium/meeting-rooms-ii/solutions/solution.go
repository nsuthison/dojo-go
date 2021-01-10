package solutions

// Question: https://leetcode.com/problems/meeting-rooms-ii/
func minMeetingRooms(intervals [][]int) int {

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
					[]TimeInterval{toBook},
				})

				break
			}

			meetingRooms[meetingRoomIdx], isAbleToBook = tryBook(toBook, meetingRooms[meetingRoomIdx])

			meetingRoomIdx++
		}

		if !isAbleToBook {
			meetingRooms = append(meetingRooms, MeetingRoom{
				[]TimeInterval{toBook},
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

type TimeInterval struct {
	Start int
	End   int
}

func (currentTimeInterval *TimeInterval) IsBefore(anotherInterval TimeInterval) bool {
	return currentTimeInterval.End <= anotherInterval.Start
}

func (currentTimeInterval *TimeInterval) IsAfter(anotherInterval TimeInterval) bool {
	return currentTimeInterval.Start >= anotherInterval.End
}

func (currentTimeInterval *TimeInterval) IsBetween(firstInterval TimeInterval, secondInterval TimeInterval) bool {
	return currentTimeInterval.IsAfter(firstInterval) && currentTimeInterval.IsBefore(secondInterval)
}

type MeetingRoom struct {
	BookedIntervals []TimeInterval
}
