package sorting_and_searching

// pointer to the array to sort
// l - index of the first element in the partition to sort
// h - index of the last element in the partition to sort
// see l and h as first to last indexes
func QuickSort(s []int, l int, h int) {
	var p int // index of partition

	// if the last and first index is still not the same;
	// meaning the partition is not yet empty
	if (h - l) > 0 {
		p = partition(s, l, h) // find the partition index
		QuickSort(s, l, p-1)
		QuickSort(s, p+1, h)
	}
}

func partition(s []int, l int, h int) int {
	var p int         // pivot element index
	var firstHigh int // divider position for pivot element

	// start with last element in the partition as pivot element
	// and first element as first high
	p = h
	firstHigh = l

	for i := l; i < h; i++ {
		if s[i] < s[p] {
			swap(&s[i], &s[firstHigh])
			firstHigh++
		}
	}
	swap(&s[p], &s[firstHigh])

	return firstHigh
}
