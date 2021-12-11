package Int64Utils

func Max(list ...int64) int64 {
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

func Min(list ...int64) int64 {
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

func If(condition bool, ifTrue, ifFalse int64) int64 {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func DefaultIfZero(val, defaultVal int64) int64 {
	if val == 0 {
		return defaultVal
	}
	return val
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

func EqualsAny(val int64, list ...int64) bool {
	for _, i := range list {
		if i == val {
			return true
		}
	}
	return false
}

// InRange 左开右闭
// [start, end)
func InRange(val, start, end int64) bool {
	return start <= val && val < end
}

// OutOfRange (-∞, start) || [end, ∞)
func OutOfRange(val, start, end int64) bool {
	return !InRange(val, start, end)
}
