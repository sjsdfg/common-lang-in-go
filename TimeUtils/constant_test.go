package TimeUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	assert.Equal(t, true, Zero.IsZero())
}
