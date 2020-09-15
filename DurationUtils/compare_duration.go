package DurationUtils

import "time"

func Max(list ...time.Duration) time.Duration {
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

func Min(list ...time.Duration) time.Duration {
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

func If(condition bool, ifTrue, ifFalse time.Duration) time.Duration {
	if condition {
		return ifTrue
	}
	return ifFalse
}
