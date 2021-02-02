package StringUtils

func Ptr(s string) *string {
	return &s
}

func GetPtrValue(s *string) string {
	if s == nil {
		return Empty
	}
	return *s
}

func GetPtrValueWithDefault(s *string, def string) string {
	if s == nil {
		return def
	}
	return *s
}
