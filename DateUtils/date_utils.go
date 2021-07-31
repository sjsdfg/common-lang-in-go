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
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
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

func GetStartOfQuarter(now time.Time) time.Time {
	quarter := GetQuarter(now)
	return time.Date(now.Year(), time.Month(3*quarter-2), 1, 0, 0, 0, 0, now.Location())
}

func GetEndOfQuarter(now time.Time) time.Time {
	quarter, day := GetQuarter(now), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	return time.Date(now.Year(), time.Month(3*quarter), day, 23, 59, 59, 999999999, now.Location())
}

func GetStartOfYear(now time.Time) time.Time {
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
}

func GetEndOfYear(now time.Time) time.Time {
	return time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, now.Location())
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

func GetQuarter(now time.Time) int {
	month := now.Month()
	return (int(month) + 2) / 3
}
