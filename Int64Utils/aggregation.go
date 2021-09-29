package Int64Utils

func Sum(list ...int64) int64 {
	if len(list) <= 0 {
		return 0
	}
	result := (int64)(0)
	for _, val := range list {
		result += val
	}
	return result
}

func Ave(list ...int64) int64 {
	if len(list) <= 0 {
		return 0
	}
	return Sum(list...) / int64(len(list))
}
