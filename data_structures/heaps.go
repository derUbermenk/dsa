package data_structures

// heaps are complete binary trees
//	that have this properties:
//
//	- The root node must either have the largest(MaxHeap) or smallest(MinHeap)
//  	keys in the tree
//  - All children must be larger (MaxHeap) or smaller (MinHeap) than the parents
//
// Heaps and complete binary trees for that matter can be represented in arrays by these following
// properties.
// With i as the index of an element in an array
//   1. 2i + 1 identifies the index of the element that is that will be the left child of the ith element
//   2. 2i + 2 identifies the index of the element that is that will be the right child of the ith element
//   3. lower bound of (i - 1)/2 is the index of the parent of a node

// orders the array in a way that the elements
// will form a max heap when turned into a complete binary tree

func MaxHeap(arr []int) {
	n := len(arr)

	if n <= 0 {
		return
	}

	for i := (n / 2) - 1; i >= 0; i-- {
		HeapifyMax(arr, n, i)
	}
}

func MinHeap(arr []int) {
	n := len(arr)

	if n <= 0 {
		return
	}

	for i := (n / 2) - 1; i >= 0; i-- {
		HeapifyMin(arr, n, i)
	}
}

func HeapifyMax(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		swap(&arr[i], &arr[largest])
		HeapifyMax(arr, n, largest)
	}
}

// orders the array in a way that the elements
// will form a min heap when turned into a complete binary tree
func HeapifyMin(arr []int, n, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		swap(&arr[i], &arr[smallest])
		HeapifyMin(arr, n, smallest)
	}
}

func swap(position1, position2 *int) {
	var temp int = *position1

	*position1 = *position2
	*position2 = temp
}
