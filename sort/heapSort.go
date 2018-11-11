package sort

func HeapSort(values []int) {

	shiftUp(values, 1)
}

func shiftUp(data []int, k int) {

	if k > 1 && data[k/2] < data[k] {
		Swap(data, k/2, k)
		k /= 2
	}
}
