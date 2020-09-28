package Int64Utils

func Ptr(i int64) *int64 {
	return &i
}

func GetPtrValue(s *int64) int64 {
	if s == nil {
		return int64(0)
	}
	return *s
}

func GetPtrValueWithDefault(s *int64, def int64) int64 {
	if s == nil {
		return def
	}
	return *s
}
