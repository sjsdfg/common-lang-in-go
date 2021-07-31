package Int64Utils

import "strconv"

func ValueOf(str string) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return result
}
