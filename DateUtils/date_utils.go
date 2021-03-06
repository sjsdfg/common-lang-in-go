package DateUtils

import (
	"time"

	"github.com/sjsdfg/common-lang-in-go/IntUtils"
)

func GetStartOfHour(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
}

func GetEndOfHour(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 0, now.Location())
}

func GetStartOfDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func GetEndOfDay(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

func GetStartOfWeek(now time.Time) time.Time {
	weekday := now.Weekday()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).
		AddDate(0, 0, IntUtils.If(weekday == time.Sunday, -6, int(time.Monday-weekday)))
}

func GetEndOfWeek(now time.Time) time.Time {
	weekStart := GetStartOfWeek(now)
	return GetEndOfDay(weekStart.AddDate(0, 0, 6))
}

func GetStartOfMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func GetEndOfMonth(now time.Time) time.Time {
	nextMonth := GetStartOfMonth(now.AddDate(0, 1, 0))
	return GetEndOfDay(nextMonth.AddDate(0, 0, -1))
}

func InRange(now, start, end time.Time) bool {
	return start.UnixNano() <= now.UnixNano() && now.UnixNano() < end.UnixNano()
}

func OutOfRange(now, start, end time.Time) bool {
	return !InRange(now, start, end)
}

func SameDay(day1, day2 time.Time) bool {
	left := time.Unix(0, day1.UnixNano())
	right := time.Unix(0, day2.UnixNano())
	return left.Year() == right.Year() &&
		left.Month() == right.Month() &&
		left.Day() == right.Day()
}
