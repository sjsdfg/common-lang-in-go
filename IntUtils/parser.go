package IntUtils

import "strconv"

func FromString(str string) int {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int(result)
}
