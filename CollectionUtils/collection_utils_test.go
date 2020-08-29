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
	for i := 0; i < 1000000; i++ {
		if len(strings) <= 0 {

		}
	}
	t.Logf("{len(strings) <= 0} cost %d nanos", timer.GetDurationInNanos())

	timer.Reset()
	for i := 0; i < 1000000; i++ {
		if IsEmpty(strings) {

		}
	}
	t.Logf("{IsEmpty(strings)} cost %d nanos", timer.GetDurationInNanos())
}
