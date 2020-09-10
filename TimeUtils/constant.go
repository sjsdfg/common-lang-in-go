package TimeUtils

import "time"

var Zero = time.Unix(0, 0)

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
	NsPerDay          = MillsPerDay * MillsPerDay
)

const (
	MinutesPerHour int64 = 60
	MinutesPerDay        = MinutesPerHour * HoursPerDay
)

const (
	HoursPerDay int64 = 24
)
