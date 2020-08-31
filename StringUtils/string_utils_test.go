package StringUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, IsEmpty(Empty))
}

func TestIsNotEmpty(t *testing.T) {
	testCase := "test"
	assert.Equal(t, true, IsNotEmpty(testCase))
}

func TestIsBlank(t *testing.T) {
	assert.Equal(t, true, IsBlank(" "))
	assert.Equal(t, true, IsBlank("	"))
}

func TestIsNotBlank(t *testing.T) {
	assert.Equal(t, true, IsNotBlank(" 123"))
	assert.Equal(t, false, IsNotBlank(" "))
}

func TestIsAllEmpty(t *testing.T) {
	assert.Equal(t, true, IsAllEmpty("", ""))
	assert.Equal(t, false, IsAllEmpty("", "123"))
}

func TestIsAnyEmpty(t *testing.T) {
	assert.Equal(t, true, IsAnyEmpty("", "123"))
	assert.Equal(t, false, IsAnyEmpty("123", "324"))
}
