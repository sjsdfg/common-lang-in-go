package Float32Utils

import "strconv"

func ValueOf(str string) float32 {
	result, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(result)
}
