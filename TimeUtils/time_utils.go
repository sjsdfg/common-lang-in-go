package TimeUtils

import (
	"time"
)

func GetCurrentMills() int64 {
	return time.Now().Unix() * MillsPerSecond
}

func GetMills(time time.Time) int64 {
	return time.UnixNano() / NsPerMill
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
