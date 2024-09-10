package merge_sort

import (
	"cmp"
)

type SortFunction[T cmp.Ordered] func(first, second T) bool

func Sort[T cmp.Ordered](items []T, sort SortFunction[T]) []T {
	length := len(items)
	if length < 2 {
		return items
	}

	mid := length / 2

	return sortAndMerge(Sort(items[:mid], sort), Sort(items[mid:], sort), sort)
}

func sortAndMerge[T cmp.Ordered](left, right []T, sort SortFunction[T]) []T {
	leftLength, rightLength := len(left), len(right)
	leftIndex, rightIndex := 0, 0
	result := make([]T, 0, leftLength+rightLength)

	for leftIndex < leftLength && rightIndex < rightLength {
		if sort(left[leftIndex], right[rightIndex]) {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	for ; leftIndex < leftLength; leftIndex++ {
		result = append(result, left[leftIndex])
	}
	for ; rightIndex < rightLength; rightIndex++ {
		result = append(result, right[rightIndex])
	}

	return result
}
