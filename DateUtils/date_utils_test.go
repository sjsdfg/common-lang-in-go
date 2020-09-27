package DateUtils

import (
	"testing"
	"time"
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
