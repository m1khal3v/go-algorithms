package merge_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	items := []uint8{1, 3, 5, 4, 2}
	assert.Equal(t, []uint8{1, 2, 3, 4, 5}, Sort(items, func(first, second uint8) bool {
		return first < second
	}))
}

func TestReverseSort(t *testing.T) {
	items := []uint8{1, 3, 5, 4, 2}
	assert.Equal(t, []uint8{5, 4, 3, 2, 1}, Sort(items, func(first, second uint8) bool {
		return first > second
	}))
}
