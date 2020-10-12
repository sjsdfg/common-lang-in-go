package IntUtils

import "math"

func Max(list ...int) int {
	if len(list) <= 0 {
		return math.MaxInt32
	}
	max := list[0]
	for _, val := range list {
		if val > max {
			max = val
		}
	}
	return max
}

func Min(list ...int) int {
	if len(list) <= 0 {
		return math.MinInt32
	}
	min := list[0]
	for _, val := range list {
		if val < min {
			min = val
		}
	}
	return min
}

func If(condition bool, ifTrue, ifFalse int) int {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func Abs(a int) int {
	return If(a > 0, a, -a)
}

func Distinct(list ...int) []int {
	if len(list) <= 0 {
		return nil
	}
	result := make([]int, 0, len(list))
	distinctMap := make(map[int]struct{}, len(list))

	for _, s := range list {
		if _, found := distinctMap[s]; !found {
			distinctMap[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}
