package quick_sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"testing"
)

func TestSort(t *testing.T) {
	items := rand.Perm(32)
	Sort(items, func(first, second int) bool {
		return first < second
	})
	assert.Equal(t, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31,
	}, items)
}

func TestReverseSort(t *testing.T) {
	items := rand.Perm(32)
	Sort(items, func(first, second int) bool {
		return first > second
	})
	assert.Equal(t, []int{
		31, 30, 29, 28, 27, 26, 25,
		24, 23, 22, 21, 20, 19, 18, 17,
		16, 15, 14, 13, 12, 11, 10, 9, 8,
		7, 6, 5, 4, 3, 2, 1, 0,
	}, items)
}

func TestSortSmall(t *testing.T) {
	items := []uint8{1, 3, 5, 4, 2}
	Sort(items, func(first, second uint8) bool {
		return first < second
	})
	assert.Equal(t, []uint8{1, 2, 3, 4, 5}, items)
}

func TestReverseSortSmall(t *testing.T) {
	items := []uint8{1, 3, 5, 4, 2}
	Sort(items, func(first, second uint8) bool {
		return first > second
	})
	assert.Equal(t, []uint8{5, 4, 3, 2, 1}, items)
}
