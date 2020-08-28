package NumberUtils

func MaxFloat64(list ...float64) float64 {
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

func MinFloat64(list ...float64) float64 {
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
