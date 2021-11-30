package Cast

type timeArg struct {
	format string
}

type TimeOption func(*timeArg)

// TimeFormat sets format to time.Parse.
func TimeFormat(format string) TimeOption {
	return func(arg *timeArg) {
		arg.format = format
	}
}
