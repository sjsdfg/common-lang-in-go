package DateUtils

import (
	"github.com/sjsdfg/common-lang-in-go/IntUtils"
	"time"
)

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
