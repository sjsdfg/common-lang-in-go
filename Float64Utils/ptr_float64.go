package Float64Utils

func Ptr(f float64) *float64 {
	return &f
}

func GetPtrValue(s *float64) float64 {
	if s == nil {
		return float64(0)
	}
	return *s
}

func GetPtrValueWithDefault(s *float64, def float64) float64 {
	if s == nil {
		return def
	}
	return *s
}
