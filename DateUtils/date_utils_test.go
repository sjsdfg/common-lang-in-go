package DateUtils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetStartOfDay(t *testing.T) {
	t.Log(GetStartOfDay(time.Now()))
}

func TestGetEndOfDay(t *testing.T) {
	t.Log(GetEndOfDay(time.Now()))
}

func TestGetStartOfWeek(t *testing.T) {
	t.Log(GetStartOfWeek(time.Now()))
}

func TestGetEndOfWeek(t *testing.T) {
	t.Log(GetEndOfWeek(time.Now()))
}

func TestGetStartOfMonth(t *testing.T) {
	t.Log(GetStartOfMonth(time.Now()))
}

func TestGetEndOfMonth(t *testing.T) {
	t.Log(GetEndOfMonth(time.Now()))
}

func TestGetQuarter(t *testing.T) {
	assert.Equal(t, 1, GetQuarter(time.Date(2021, 01, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 1, GetQuarter(time.Date(2021, 02, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 1, GetQuarter(time.Date(2021, 03, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 2, GetQuarter(time.Date(2021, 04, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 2, GetQuarter(time.Date(2021, 05, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 2, GetQuarter(time.Date(2021, 06, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 3, GetQuarter(time.Date(2021, 07, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 3, GetQuarter(time.Date(2021, 8, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 3, GetQuarter(time.Date(2021, 9, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 4, GetQuarter(time.Date(2021, 10, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 4, GetQuarter(time.Date(2021, 11, 05, 12, 12, 0, 0, time.Local)))
	assert.Equal(t, 4, GetQuarter(time.Date(2021, 12, 05, 12, 12, 0, 0, time.Local)))
}

func TestGetStartOfQuarter(t *testing.T) {
	t.Log(GetStartOfQuarter(time.Date(2021, 1, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 1, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 2, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 2, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 3, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 3, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 4, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 4, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 5, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 5, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 6, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 6, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 7, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 7, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 8, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 8, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 9, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 9, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 10, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 10, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 11, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 11, 05, 12, 12, 0, 0, time.Local)))

	t.Log(GetStartOfQuarter(time.Date(2021, 12, 05, 12, 12, 0, 0, time.Local)))
	t.Log(GetEndOfQuarter(time.Date(2021, 12, 05, 12, 12, 0, 0, time.Local)))
}

func TestGetStartOfYear(t *testing.T) {
	t.Log(GetStartOfYear(time.Now()))
}

func TestGetEndOfYear(t *testing.T) {
	t.Log(GetEndOfYear(time.Now()))
}
