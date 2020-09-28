package Float32Utils

func Ptr(f float32) *float32 {
	return &f
}

func GetPtrValue(s *float32) float32 {
	if s == nil {
		return float32(0)
	}
	return *s
}

func GetPtrValueWithDefault(s *float32, def float32) float32 {
	if s == nil {
		return def
	}
	return *s
}
