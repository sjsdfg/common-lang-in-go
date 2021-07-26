package StringUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, IsEmpty(""))
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

func TestDefaultIfZero(t *testing.T) {
	var str = "123"
	assert.Equal(t, str, DefaultIfZero("0", str))
	assert.Equal(t, str, DefaultIfZero(str, ""))
}

func TestDefaultIf(t *testing.T) {
	var str = "123"
	assert.Equal(t, str, DefaultIf(IsZero, "0", str))
	assert.Equal(t, str, DefaultIf(IsEmpty, str, ""))
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

func TestTruncate(t *testing.T) {
	testCase := "0123456789"
	assert.Equal(t, "012", Truncate(testCase, 0, 3))
	assert.Equal(t, "012", Truncate(testCase, -5, 3))
	assert.Equal(t, testCase, Truncate(testCase, -5, 20))
}

func TestDistinct(t *testing.T) {
	testCase := []string{"1", "1", "1"}
	testCase2 := []string{"2", "2", "2"}
	var testCase3 []string
	assert.Equal(t, []string{"1"}, Distinct(testCase...))
	assert.Equal(t, []string{"1", "2"}, Distinct(append(testCase, testCase2...)...))
	assert.Equal(t, []string(nil), Distinct(testCase3...))
}

func TestReverse(t *testing.T) {
	testCase := "123456789"
	result := "987654321"
	assert.Equal(t, result, Reverse(testCase))
}

func TestReplaceHolder(t *testing.T) {
	str := "{holder} must be {replace}, other {holder} same"
	holderMap := map[string]string{
		"{holder}":  "liming",
		"{replace}": "clever",
	}
	assert.Equal(t,
		"liming must be clever, other liming same",
		ReplaceHolder(str, holderMap))
}
