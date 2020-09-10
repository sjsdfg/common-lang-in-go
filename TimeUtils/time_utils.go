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

func MaxDuration(list ...time.Duration) time.Duration {
	if len(list) <= 0 {
		return 0
	}
	max := list[0]
	for _, t := range list {
		if t > max {
			max = t
		}
	}
	return max
}

func MinDuration(list ...time.Duration) time.Duration {
	if len(list) <= 0 {
		return 0
	}
	min := list[0]
	for _, t := range list {
		if t < min {
			min = t
		}
	}
	return min
}

func IfDuration(condition bool, ifTrue, ifFalse time.Duration) time.Duration {
	if condition {
		return ifTrue
	}
	return ifFalse
}
