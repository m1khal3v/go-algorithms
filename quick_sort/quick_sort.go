package quick_sort

import (
	"cmp"
	"math/rand/v2"
	"sync"
)

type SortFunction[T cmp.Ordered] func(first, second T) bool

func Sort[T cmp.Ordered](items []T, sort SortFunction[T]) {
	quickSort(items, 0, len(items)-1, sort)
}

func quickSort[T cmp.Ordered](items []T, low, high int, sort SortFunction[T]) {
	length := high - low + 1
	if length < 2 {
		return
	}

	// choose random pivot and move it to the end
	pivotIndex := rand.IntN(length) + low
	pivot := items[pivotIndex]
	items[pivotIndex], items[high] = items[high], items[pivotIndex]

	// move elements less than pivot to the left
	for index := low; index < high; index++ {
		if sort(items[index], pivot) {
			items[index], items[low] = items[low], items[index]
			low++
		}
	}

	// move pivot after last smaller item
	items[low], items[high] = items[high], items[low]

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		quickSort(items, 0, low-1, sort)
	}()
	go func() {
		defer wg.Done()
		quickSort(items, low+1, high, sort)
	}()
	wg.Wait()
}
