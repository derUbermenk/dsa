package sorting_and_searching

// swaps position 1 and position2 in s
func swap(position1, position2 *int) {
	temp := *position1
	*position1 = *position2
	*position2 = temp
}

// deqeues the last element in the array
func deqeue(arr []int) int {
	last_indx := len(arr) - 1

	last_element := arr[last_indx]
	arr = arr[:last_indx]

	return last_element
}
