package BoolUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnd(t *testing.T) {
	assert.Equal(t, true, And(true, true, true))
	assert.Equal(t, false, And(false, true))
}

func TestOr(t *testing.T) {
	assert.Equal(t, true, Or(false, false, true))
	assert.Equal(t, false, Or(false, false, false))
}

func TestXor(t *testing.T) {
	assert.Equal(t, true, Xor(true, false))
	assert.Equal(t, false, Xor(false, false))
	assert.Equal(t, false, Xor(true, false, false, true))
	assert.Equal(t, true, Xor(true, false, true, false, true, false))
}
