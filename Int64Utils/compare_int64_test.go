package Int64Utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsAny(t *testing.T) {
	ints := []int64{1, 5, 7}
	assert.Equal(t, true, EqualsAny(1, append(ints, 5)...))

	t.Log(ints)
}
