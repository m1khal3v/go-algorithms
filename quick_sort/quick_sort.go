package quick_sort

import (
	"cmp"
	"math/rand/v2"
)

type SortFunction[T cmp.Ordered] func(first, second T) bool

func Sort[T cmp.Ordered](items []T, sort SortFunction[T]) {
	quickSort(items, 0, len(items)-1, sort)
}

func quickSort[T cmp.Ordered](items []T, low, high int, sort SortFunction[T]) {
	const minLength = 20

	length := high - low + 1
	if length < 2 {
		return
	}

	// if the slice is small, then insertion sort is used as an optimization
	if length < minLength {
		insertionSort(items, low, high, sort)
		return
	}

	mid := low

	// choose random pivot and move it to the end
	pivotIndex := rand.IntN(length) + low
	pivot := items[pivotIndex]
	items[pivotIndex], items[high] = items[high], items[pivotIndex]

	// move elements less than pivot to the left
	for index := low; index < high; index++ {
		if sort(items[index], pivot) {
			items[index], items[mid] = items[mid], items[index]
			mid++
		}
	}

	// move pivot after last smaller item
	items[mid], items[high] = items[high], items[mid]

	quickSort(items, low, mid-1, sort)
	quickSort(items, mid+1, high, sort)

	return
}

func insertionSort[T cmp.Ordered](items []T, low, high int, sort SortFunction[T]) {
	for i := low; i <= high; i++ {
		for j := i; j > low && sort(items[j], items[j-1]); j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
}
