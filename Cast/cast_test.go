package Cast

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	num := 5
	assert.Equal(t, ToString(1), fmt.Sprint(1))
	t.Log(ToString(1), fmt.Sprint(1))
	assert.Equal(t, ToString(20.5), fmt.Sprint(20.5))
	t.Log(ToString(20.5), fmt.Sprint(20.5))
	assert.Equal(t, ToString(&num), fmt.Sprint(num))
	t.Log(ToString(&num), fmt.Sprint(num))
}

func BenchmarkName(b *testing.B) {
	num := 5
	b.Run("Cast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToString(1)
			ToString(20.5)
			ToString(&num)
		}
	})

	b.Run("fmt.Sprint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Sprint(1)
			fmt.Sprint(20.5)
			fmt.Sprint(num)
		}
	})
}
