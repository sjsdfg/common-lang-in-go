package Float32Utils

func Max(list ...float32) float32 {
	if len(list) <= 0 {
		return 0
	}
	max := list[0]
	for _, val := range list {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(list ...float32) float32 {
	if len(list) <= 0 {
		return 0
	}
	min := list[0]
	for _, val := range list {
		if val < min {
			min = val
		}
	}
	return min
}

func Condition(condition bool, ifTrue, ifFalse float32) float32 {
	if condition {
		return ifTrue
	}
	return ifFalse
}