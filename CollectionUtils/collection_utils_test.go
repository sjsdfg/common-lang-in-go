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

	assert.Equal(t, true, IsEmpty([]string(nil)))
	assert.Equal(t, true, IsEmpty(nil))
	assert.Equal(t, false, IsEmpty(2))
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

func BenchmarkIsEmpty(b *testing.B) {
	strings := make([]string, 0)
	for i := 0; i < b.N; i++ {
		if IsEmpty(strings) {

		}
	}
}

func BenchmarkNativeIsEmpty(b *testing.B) {
	strings := make([]string, 0)
	for i := 0; i < b.N; i++ {
		if IsEmpty(strings) {

		}
	}
}

func TestForEach(t *testing.T) {
	students := createStudents()
	ForEach(students, func(index int) {
		t.Log(students[index])
	})

	ForEach(students, func(index int) {
		students[index].name = "testing"
	})

	ForEach(students, func(index int) {
		assert.Equal(t, "testing", students[index].name)
	})
}
