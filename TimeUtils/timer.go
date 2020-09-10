package TimeUtils

import (
	"time"
)

type Timer struct {
	startTime time.Time
}

func NewTimer() Timer {
	return Timer{startTime: time.Now()}
}

func (t *Timer) GetDurationInSeconds() int64 {
	return time.Now().Unix() - t.startTime.Unix()
}

func (t *Timer) GetDurationInMills() int64 {
	return t.GetDurationInNanos() / NsPerMill
}

func (t *Timer) GetDurationInNanos() int64 {
	return time.Now().UnixNano() - t.startTime.UnixNano()
}

func (t *Timer) Reset() {
	t.startTime = time.Now()
}
