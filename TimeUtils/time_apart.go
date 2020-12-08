package TimeUtils

import (
	"github.com/sjsdfg/common-lang-in-go/Int64Utils"
	"time"
)

func DaysApart(t1, t2 time.Time) int64 {
	t1Local, t2Local := t1.Local(), t2.Local()

	startDayOfT1 := time.Date(t1Local.Year(), t1Local.Month(), t1Local.Day(), 0, 0, 0, 0, time.Local)
	startDayOfT2 := time.Date(t2Local.Year(), t2Local.Month(), t2Local.Day(), 0, 0, 0, 0, time.Local)

	return int64(startDayOfT1.Sub(startDayOfT2).Hours()) / HoursPerDay
}

func AbsDaysApart(t1, t2 time.Time) int64 {
	return Int64Utils.Abs(DaysApart(t1, t2))
}

func MinutesApart(t1, t2 time.Time) int64 {
	return int64(t1.Local().Sub(t2.Local()).Minutes())
}

func AbsMinutesApart(t1, t2 time.Time) int64 {
	return Int64Utils.Abs(MinutesApart(t1, t2))
}

func HoursApart(t1, t2 time.Time) int64 {
	return int64(t1.Local().Sub(t2.Local()).Hours())
}

func AbsHoursApart(t1, t2 time.Time) int64 {
	return Int64Utils.Abs(HoursApart(t1, t2))
}
