package TimeUtils

import (
	"github.com/sjsdfg/common-lang-in-go/DateUtils"
	"time"
)

func GetCurrentMills() int64 {
	return time.Now().Unix() * DateUtils.MillsPerSeconds
}

func GetMills(time time.Time) int64 {
	return time.Unix() * DateUtils.MillsPerSeconds
}
