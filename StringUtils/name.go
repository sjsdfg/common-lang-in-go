package StringUtils

func TruncateName(name string, saveLength int) string {
	runes := []rune(name)
	if len(runes) <= saveLength {
		return name
	}
	return string(runes[:saveLength]) + "..."
}
