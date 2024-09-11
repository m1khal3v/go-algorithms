package binary_search

import (
	"math/rand/v2"
	"testing"
)

const count = 10000000

func BenchmarkFind(b *testing.B) {
	b.StopTimer()
	items := make([]int, count)
	for i := 0; i < count; i++ {
		items[i] = i
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Find(items, rand.IntN(count))
	}
}
