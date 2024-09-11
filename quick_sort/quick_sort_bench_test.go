package quick_sort

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		items := rand.Perm(1024)
		b.StartTimer()

		Sort(items, func(first, second int) bool {
			return first < second
		})
	}
}
