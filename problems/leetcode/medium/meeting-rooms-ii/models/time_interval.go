package models

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
