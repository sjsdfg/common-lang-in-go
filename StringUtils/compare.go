package StringUtils

// DigitalBiggerThan 假设：入参 a 和 b 都是合法的正整数类型
// 并且没有多余的前导 0
func DigitalBiggerThan(a, b string) bool {
	lenA := len(a)
	lenB := len(b)
	if lenA == lenB {
		// 长度相等根据字符串判断
		return a > b
	}
	// 不等根据长度判断
	return lenA > lenB
}
