package Float64Utils

import "strconv"

func ValueOf(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}
