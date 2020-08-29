package DateUtils

const (
	MillsPerSeconds int64 = 1000
	MillsPerMinute  int64 = MillsPerSeconds * 60
	MillsPerHour    int64 = MillsPerMinute * 60
	MillsPerDay     int64 = MillsPerHour * 24
)

const (
	NanosPerMills   int64 = 1e6
	NanosPerSeconds int64 = MillsPerSeconds * NanosPerMills
	NanosPerMinute  int64 = MillsPerMinute * NanosPerMills
	NanosPerHour    int64 = MillsPerHour * NanosPerMills
	NanosPerDay     int64 = MillsPerDay * MillsPerDay
)
