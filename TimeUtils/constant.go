package TimeUtils

import "time"

var Zero = time.Unix(0, 0)

const (
	MillsPerSecond int64 = 1000
	MillsPerMinute int64 = MillsPerSecond * 60
	MillsPerHour   int64 = MillsPerMinute * 60
	MillsPerDay    int64 = MillsPerHour * 24
)

const (
	NsPerMill   int64 = 1e6
	NsPerSecond int64 = MillsPerSecond * NsPerMill
	NsPerMinute int64 = MillsPerMinute * NsPerMill
	NsPerHour   int64 = MillsPerHour * NsPerMill
	NsPerDay    int64 = MillsPerDay * MillsPerDay
)
