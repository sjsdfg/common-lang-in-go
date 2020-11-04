package Int64Utils

import (
	"strconv"
)

// 每隔 3 位，加一个单字符的逗号 ,
func FormatToFinancial(number int64) string {
	str := strconv.FormatInt(number, 10)
	if len(str) <= 3 {
		return str
	}
	if number > 0 {
		return comma(str)
	} else {
		return "-" + comma(str[1:])
	}
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3]) + "," + comma(s[len(s)-3:])
}
