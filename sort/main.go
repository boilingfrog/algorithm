package main

import "fmt"

func main() {
	var arr = []int{1000, 11, 3, 100, 100, 5, 6, 7, 8, 9, 4, 1}
	//fmt.Println(data)
	////MaxHeap(arr, len(arr))
	//HeapSort(arr, len(arr))
	//fmt.Println(data)

	// 快排
	//QuickSort(arr, 9)
	bubbleInsert(arr)
	fmt.Println(arr)
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
