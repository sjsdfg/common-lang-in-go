package CollectionUtils

import (
	"fmt"
	"github.com/sjsdfg/common-lang-in-go/TimeUtils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type student struct {
	id     int
	name   string
	age    int
	scores []int
}

func (s *student) String() string {
	return fmt.Sprintf("{id:%d, name:%s, age:%d,scores:%v}", s.id, s.name, s.age, s.scores)
}

func createStudents() []*student {
	names := []string{"Tom", "Kate", "Lucy", "Jim", "Jack", "King", "Lee", "Mask"}
	students := make([]*student, 10)
	rnd := func(start, end int) int { return rand.Intn(end-start) + start }
	for i := 0; i < 10; i++ {
		students[i] = &student{
			id:     i + 1,
			name:   names[rand.Intn(len(names))],
			age:    rnd(15, 26),
			scores: []int{rnd(60, 100), rnd(60, 100), rnd(60, 100)},
		}
	}
	return students
}

func TestMapToStringSlice(t *testing.T) {
	students := createStudents()
	timer := TimeUtils.NewTimer()
	reflectSlice := MapToStringSlice(students, func(i interface{}) string {
		return i.(*student).name
	})
	t.Logf("MapToStringSlice cost %d nanos", timer.GetDurationInNanos())

	timer.Reset()
	nativeSlice := NativeMapToStringSlice(students)
	t.Logf("NativeMapToStringSlice cost %d nanos", timer.GetDurationInNanos())
	assert.Equal(t, nativeSlice, reflectSlice)
}

func NativeMapToStringSlice(list []*student) []string {
	if len(list) <= 0 {
		return []string{}
	}
	result := make([]string, 0, len(list))
	for _, s := range list {
		result = append(result, s.name)
	}
	return result
}

func BenchmarkMapToStringSlice(b *testing.B) {
	students := createStudents()
	for i := 0; i < 1000; i++ {
		MapToStringSlice(students, func(i interface{}) string {
			return i.(*student).name
		})
	}
}

func BenchmarkNativeMapToStringSlice(b *testing.B) {
	students := createStudents()
	for i := 0; i < 1000; i++ {
		NativeMapToStringSlice(students)
	}
}

func TestMapToIntSlice(t *testing.T) {
	students := createStudents()
	ageSlice := MapToIntSlice(students, func(i interface{}) int {
		return i.(*student).age
	})
	t.Log(ageSlice)
}
