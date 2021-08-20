package TimeUtils

import (
	"time"
)

func GetCurrentMills() int64 {
	return GetMills(time.Now())
}

func GetMills(time time.Time) int64 {
	return time.UnixNano() / NsPerMill
}

func FromMills(mills int64) time.Time {
	return time.Unix(mills/MillsPerSecond, 0)
}

func Max(list ...time.Time) time.Time {
	if len(list) <= 0 {
		return Zero
	}
	max := list[0]
	for _, t := range list {
		if t.After(max) {
			max = t
		}
	}
	return max
}

func Min(list ...time.Time) time.Time {
	if len(list) <= 0 {
		return Zero
	}
	min := list[0]
	for _, t := range list {
		if t.Before(min) {
			min = t
		}
	}
	return min
}

func If(condition bool, ifTrue, ifFalse time.Time) time.Time {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func DefaultIfZero(val, defaultVal time.Time) time.Time {
	if val.IsZero() {
		return defaultVal
	}
	return val
}
