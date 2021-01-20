package models

// TimeInterval represent any duration of time
type TimeInterval struct {
	Start int
	End   int
}

// IsBefore check that is the given anotherInterval is before current time object
func (currentTimeInterval *TimeInterval) IsBefore(anotherInterval TimeInterval) bool {
	return currentTimeInterval.End <= anotherInterval.Start
}

// IsAfter check that is the given anotherInterval is after current time object
func (currentTimeInterval *TimeInterval) IsAfter(anotherInterval TimeInterval) bool {
	return currentTimeInterval.Start >= anotherInterval.End
}

// IsBetween check that is the given anotherInterval is in between current time object
func (currentTimeInterval *TimeInterval) IsBetween(firstInterval TimeInterval, secondInterval TimeInterval) bool {
	return currentTimeInterval.IsAfter(firstInterval) && currentTimeInterval.IsBefore(secondInterval)
}
