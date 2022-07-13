package sorting_and_searching

import (
	"dsa/data_structures"
)

func HeapSort(arr []int) {
	max_heap := append([]int{}, arr...)
	data_structures.MaxHeap(max_heap)

	// in the max heap, the root is the largest number
	// steps are:
	// 	1. swap with left most node; last index in the max heap array
	//  2. remove the last element in the max heap array (this is the current largest due to swap)
	//  3. heapify (this puts the next largest up top)
	//  4. swap with last node
	//  5. do this until max heap is empty

	/*
		while max_heap not empty:
			swap(max_heap[0], max_heap[last_element])
			s[i] = deqeue(max_heap)
			heapify(max_heap)
	*/

	for i := len(arr) - 1; i >= 0; i-- {
		swap(&max_heap[0], &max_heap[i])

		arr[i] = max_heap[i]
		max_heap = max_heap[:i]

		data_structures.HeapifyMax(max_heap, i, 0)
	}
}
