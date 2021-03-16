package StringUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitalBiggerThan(t *testing.T) {
	assert.Equal(t, false, DigitalBiggerThan("1", "1"))
	assert.Equal(t, true, DigitalBiggerThan("2", "1"))
	assert.Equal(t, true, DigitalBiggerThan("21", "4"))
}
