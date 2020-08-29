package TimeUtils

import (
	"github.com/sjsdfg/common-lang-in-go/DateUtils"
	"time"
)

type Timer struct {
	startTime time.Time
}

func NewTimer() Timer {
	return Timer{startTime: time.Now()}
}

func (t Timer) GetDurationInSeconds() int64 {
	return time.Now().Unix() - t.startTime.Unix()
}

func (t Timer) GetDurationInMills() int64 {
	return t.GetDurationInSeconds() * DateUtils.MillsPerSeconds
}

func (t Timer) GetDurationInNanos() int64 {
	return t.GetDurationInMills() * DateUtils.NanosPerMills
}

func (t Timer) Reset() {
	t.startTime = time.Now()
}
