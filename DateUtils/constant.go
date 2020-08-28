package DateUtils

const (
	MillsPerSeconds int64 = 1000
	MillsPerMinute  int64 = MillsPerSeconds * 60
	MillsPerHour    int64 = MillsPerMinute * 60
	MillsPerDay     int64 = MillsPerHour * 24
)

const (
	NanosPerSeconds int64 = 1e6
	NanosPerMinute  int64 = MillsPerMinute * NanosPerSeconds
	NanosPerHour    int64 = MillsPerHour * NanosPerSeconds
	NanosPerDay     int64 = MillsPerDay * MillsPerDay
)
