package main

import "fmt"

func main() {
	var ar = []int{1, 2, 2, 2, 2, 3, 4, 4, 5, 6, 7, 8, 9, 10, 1999}
	//
	//// 快排
	////QuickSort(arr, 9)
	//countingSort(ar)
	//fmt.Println(ar)

	fmt.Println(bsearchLastEqual(ar, 4))
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
