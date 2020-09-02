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

func TestIsAnyNoneEmpty(t *testing.T) {
	assert.Equal(t, true, IsAnyNoneEmpty("", "123"))
	assert.Equal(t, false, IsAnyNoneEmpty())
	assert.Equal(t, false, IsAnyNoneEmpty("", ""))
}

func TestEqual(t *testing.T) {
	assert.Equal(t, true, Equal("123", "123"))
	assert.Equal(t, true, Equal("abc", "abc"))
	assert.Equal(t, false, Equal("abc", "Abc"))
}

func TestEqualIgnoreCase(t *testing.T) {
	assert.Equal(t, true, EqualIgnoreCase("abc", "abc"))
	assert.Equal(t, true, EqualIgnoreCase("abc", "Abc"))
}

func TestEqualsAny(t *testing.T) {
	assert.Equal(t, true, EqualsAny("123", "123", "345", "abc"))
	assert.Equal(t, false, EqualsAny("123", "345", "abc"))
	assert.Equal(t, false, EqualsAny("abc", "345", "Abc"))
}

func TestEqualsAnyIgnoreCase(t *testing.T) {
	assert.Equal(t, true, EqualsAnyIgnoreCase("123", "123", "345", "abc"))
	assert.Equal(t, false, EqualsAnyIgnoreCase("123", "345", "abc"))
	assert.Equal(t, true, EqualsAnyIgnoreCase("abc", "345", "Abc"))
}

func TestIsDigital(t *testing.T) {
	assert.Equal(t, true, IsDigital("123456"))
	assert.Equal(t, false, IsDigital("123asdasd"))
}

func TestDefaultIfEmpty(t *testing.T) {
	var str = "123"
	assert.Equal(t, str, DefaultIfEmpty("", str))
	assert.Equal(t, str, DefaultIfEmpty(str, ""))
}

func TestIsZero(t *testing.T) {
	assert.Equal(t, true, IsZero(Empty))
	assert.Equal(t, true, IsZero("0"))
	assert.Equal(t, false, IsZero("2"))
}

func TestIsAllZero(t *testing.T) {
	assert.Equal(t, true, IsAllZero("", ""))
	assert.Equal(t, true, IsAllZero("", "0"))
	assert.Equal(t, false, IsAllZero("", "123"))
}
