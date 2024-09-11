package binary_search

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestFind(t *testing.T) {
	items := rand.Perm(1024)
	slices.Sort(items)

	index := rand.IntN(len(items))
	target := items[index]

	result, ok := Find(items, target)
	require.True(t, ok)
	assert.Equal(t, index, result)
}

func TestFindEmpty(t *testing.T) {
	result, ok := Find([]int{}, rand.Int())
	require.False(t, ok)
	assert.Equal(t, 0, result)
}

func TestFindNotFound(t *testing.T) {
	result, ok := Find([]int{-1, 0, 1, 2, 3, 4, 5, 6}, rand.IntN(7)+7)
	require.False(t, ok)
	assert.Equal(t, 0, result)
}
