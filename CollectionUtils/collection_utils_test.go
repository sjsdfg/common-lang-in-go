package CollectionUtils

import (
	"github.com/sjsdfg/common-lang-in-go/TimeUtils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	strings := make([]string, 0)
	assert.Equal(t, len(strings) <= 0, IsEmpty(strings))

	m := make(map[int]int, 5)
	assert.Equal(t, len(m) <= 0, IsEmpty(m))

	i := make([]string, 0, 20)
	assert.Equal(t, len(i) <= 0, IsEmpty(i))

	strings = make([]string, 2)
	assert.Equal(t, len(strings) <= 0, IsEmpty(strings))
}

func TestIsEmptySpeed(t *testing.T) {
	strings := make([]string, 0)
	timer := TimeUtils.NewTimer()
	retryTime := 100_000_000_000
	for i := 0; i < retryTime; i++ {
		if len(strings) <= 0 {

		}
	}

	nativeCost := timer.GetDurationInNanos()
	t.Logf("{len(strings) <= 0} cost %d nanos", nativeCost)

	timer.Reset()
	for i := 0; i < retryTime; i++ {
		if IsEmpty(strings) {

		}
	}
	reflectCost := timer.GetDurationInNanos()
	t.Logf("{IsEmpty(strings)} cost %d nanos", reflectCost)
	t.Logf("reflectCost / nativeCost = %d", reflectCost/nativeCost)
}
