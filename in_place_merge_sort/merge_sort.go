package merge_sort

import (
	"cmp"
)

type SortFunction[T cmp.Ordered] func(first, second T) bool

func Sort[T cmp.Ordered](items []T, sort SortFunction[T]) {
	mergeSort(items, 0, len(items)-1, sort)
}

func mergeSort[T cmp.Ordered](items []T, low, high int, sort SortFunction[T]) {
	if low >= high {
		return
	}

	mid := (high-low)/2 + low

	mergeSort(items, low, mid, sort)
	mergeSort(items, mid+1, high, sort)

	sortAndMerge(items, low, mid, high, sort)
}

func sortAndMerge[T cmp.Ordered](items []T, low, mid, high int, sort SortFunction[T]) {
	lowSecond := mid + 1

	if sort(items[mid], items[lowSecond]) {
		return
	}

	for low <= mid && lowSecond <= high {
		if sort(items[low], items[lowSecond]) {
			low++
		} else {
			lowSecondValue := items[lowSecond]

			for index := lowSecond; index != low; index-- {
				items[index] = items[index-1]
			}

			items[low] = lowSecondValue

			low++
			lowSecond++
			mid++
		}
	}
}
