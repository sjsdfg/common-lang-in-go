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

func If(condition bool, ifTrue, ifFalse float32) float32 {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func Abs(a float32) float32 {
	return If(a > 0, a, -a)
}

func Equal(a, b float32) bool {
	return Abs(a-b) <= 1e-6
}

// InRange 左开右闭
// [start, end)
func InRange(val, start, end float32) bool {
	return start <= val && val < end
}

// OutOfRange (-∞, start) || [end, ∞)
func OutOfRange(val, start, end float32) bool {
	return !InRange(val, start, end)
}
