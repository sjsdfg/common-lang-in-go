package Int64Utils

func Max(list ...int64) int64 {
	if len(list) <= 0 {
		return MAX
	}
	max := list[0]
	for _, val := range list {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(list ...int64) int64 {
	if len(list) <= 0 {
		return MIN
	}
	min := list[0]
	for _, val := range list {
		if val < min {
			min = val
		}
	}
	return min
}

func If(condition bool, ifTrue, ifFalse int64) int64 {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func Abs(a int64) int64 {
	return If(a > 0, a, -a)
}

func Distinct(list ...int64) []int64 {
	if len(list) <= 0 {
		return nil
	}
	result := make([]int64, 0, len(list))
	distinctMap := make(map[int64]struct{}, len(list))

	for _, s := range list {
		if _, found := distinctMap[s]; !found {
			distinctMap[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}
