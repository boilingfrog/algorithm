package main

import "fmt"

func main() {
	var arr = []int{1000, 11, 3, 5, 100, 6, 7, 8, 9, 4}
	//fmt.Println(data)
	////MaxHeap(arr, len(arr))
	//HeapSort(arr, len(arr))
	//fmt.Println(data)

	// 快排
	//QuickSort(arr, 9)
	quickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func partition1(a []int, lo, hi int) int {
	pivot := a[lo]
	i := lo - 1
	for j := lo + 1; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort1(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition1(a, lo, hi)
	quickSort1(a, lo, p-1)
	quickSort1(a, p+1, hi)
}

// 来个冒泡玩一下
func bubbleInsert(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
