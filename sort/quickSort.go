package main

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(data []int, l, u int) {
	if l < u {
		m := partition(data, l, u)
		quickSort(data, l, m-1)
		quickSort(data, m, u)
	}
}

func partition(data []int, l, u int) int {
	quick := data[l]
	left := l

	for i := l + 1; i <= u; i++ {
		if data[i] <= quick {
			left++
			data[left], data[i] = data[i], data[left]
		}
	}
	data[l], data[left] = data[left], data[l]
	return left + 1
}
