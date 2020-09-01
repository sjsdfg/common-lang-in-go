package TimeUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := NewTimer()
	time.Sleep(1 * time.Second)
	assert.Equal(t, int64(1), timer.GetDurationInSeconds())
}

func TestZero(t *testing.T) {
	assert.Equal(t, int64(0), Zero.Unix())
	assert.Equal(t, int64(0), Zero.UnixNano())
}
