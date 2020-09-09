package CollectionUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllMatch(t *testing.T) {
	testCase := []string{"a", "a", "a"}
	assert.Equal(t, true, AllMatch(testCase, func(index int) bool {
		return testCase[index] == "a"
	}))

	testCase = []string{"a", "a", "b"}
	assert.Equal(t, false, AllMatch(testCase, func(index int) bool {
		return testCase[index] == "a"
	}))
}

func TestAnyMatch(t *testing.T) {
	testCase := []string{"a", "a", "b"}
	assert.Equal(t, true, AnyMatch(testCase, func(index int) bool {
		return testCase[index] == "b"
	}))

	testCase = []string{"a", "a", "b"}
	assert.Equal(t, false, AnyMatch(testCase, func(index int) bool {
		return testCase[index] == "c"
	}))
}

func TestNoneMatch(t *testing.T) {
	testCase := []string{"a", "a", "b"}
	assert.Equal(t, true, NoneMatch(testCase, func(index int) bool {
		return testCase[index] == "c"
	}))

	testCase = []string{"a", "a", "b"}
	assert.Equal(t, false, NoneMatch(testCase, func(index int) bool {
		return testCase[index] == "b"
	}))
}
