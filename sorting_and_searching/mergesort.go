package sorting_and_searching

func MergeSort(s []int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		MergeSort(s, low, middle)
		MergeSort(s, middle+1, high)
		merge(s, low, middle, high)
	}
}

func merge(s []int, low, middle, high int) {
	// make buffers that hold the Lower and Upper half of the subarrays
	L := make([]int, 0)
	H := make([]int, 0)

	// populate the buffers with the elements of the sub arrays
	for i := low; i <= middle; i++ {
		L = append(L, s[i])
	}
	for j := middle + 1; j <= high; j++ {
		H = append(H, s[j])
	}

	// iterate over both queue until either is empty
	i := low
	for len(L) != 0 && len(H) != 0 {
		if L[0] < H[0] {
			s[i] = L[0]

			// pop first element of off L
			L = L[1:]
		} else {
			s[i] = H[0]

			// pop first element of off H
			H = H[1:]
		}
		i++
	}

	for len(L) != 0 {
		s[i] = L[0]

		L = L[1:]
		i++
	}
	for len(H) != 0 {
		s[i] = H[0]

		H = H[1:]
		i++
	}
}
