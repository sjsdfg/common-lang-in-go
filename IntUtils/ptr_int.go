package IntUtils

func Ptr(i int) *int {
	return &i
}

func GetPtrValue(s *int) int {
	if s == nil {
		return int(0)
	}
	return *s
}

func GetPtrValueWithDefault(s *int, def int) int {
	if s == nil {
		return def
	}
	return *s
}
