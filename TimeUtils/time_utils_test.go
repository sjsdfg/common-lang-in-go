package TimeUtils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	now := time.Now()
	after100 := now.AddDate(100, 0, 0)
	assert.Equal(t, after100, Max(Zero, now, after100))
}

func TestMin(t *testing.T) {
	now := time.Now()
	after100 := now.AddDate(100, 0, 0)
	assert.Equal(t, Zero, Min(Zero, now, after100))
}

func TestFromMills(t *testing.T) {
	now := time.Now()
	mills := GetMills(now)
	fromMills := FromMills(mills)
	assert.Equal(t, now.Unix(), fromMills.Unix())
}
