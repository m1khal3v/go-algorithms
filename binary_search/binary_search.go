package binary_search

import (
	"cmp"
)

func Find[T cmp.Ordered](items []T, target T) (int, bool) {
	low := 0
	high := len(items) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := items[mid]

		switch {
		case guess == target:
			return mid, true
		case guess > target:
			high = mid - 1
		default: // guess < target
			low = mid + 1
		}
	}

	return 0, false
}
