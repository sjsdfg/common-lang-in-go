package Cast

import (
	"fmt"
	"testing"
)

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
