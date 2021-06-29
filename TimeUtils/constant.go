package TimeUtils

import "time"

var Zero = time.Time{}

const (
	MillsPerSecond int64 = 1000
	MillsPerMinute       = MillsPerSecond * 60
	MillsPerHour         = MillsPerMinute * 60
	MillsPerDay          = MillsPerHour * 24
)

const (
	NsPerMill   int64 = 1e6
	NsPerSecond       = MillsPerSecond * NsPerMill
	NsPerMinute       = MillsPerMinute * NsPerMill
	NsPerHour         = MillsPerHour * NsPerMill
	NsPerDay          = MillsPerDay * NsPerMill
)

const (
	SecondsPerMinutes int64 = 60
	SecondsPerHour          = SecondsPerMinutes * 60
	SecondsPerDay           = SecondsPerHour * 24
	SecondsPerWeek          = SecondsPerDay * 7
)

const (
	MinutesPerHour int64 = 60
	MinutesPerDay        = MinutesPerHour * 24
	MinutesPerWeek       = MinutesPerDay * 7
)

const (
	HoursPerDay int64 = 24
)

const (
	DaysPerWeek int64 = 7
)
