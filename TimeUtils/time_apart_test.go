package TimeUtils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDaysApart(t *testing.T) {
	testCase1 := []time.Time{
		time.Date(2020, time.December, 8, 20, 35, 0, 0, time.Local),
		time.Date(2020, time.December, 15, 10, 35, 0, 0, time.Local),
	}

	assert.Equal(t, int64(-7), DaysApart(testCase1[0], testCase1[1]))
	assert.Equal(t, int64(7), DaysApart(testCase1[1], testCase1[0]))
}

func TestHoursApart(t *testing.T) {
	utcTime := time.Date(2020, time.December, 8, 20, 35, 0, 0, time.UTC)
	localTime := time.Date(2020, time.December, 8, 20, 35, 0, 0, time.Local)

	assert.EqualValues(t, 8, HoursApart(utcTime, localTime))
	assert.EqualValues(t, -8, HoursApart(localTime, utcTime))
}
