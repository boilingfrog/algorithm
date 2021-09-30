package main

import "fmt"

func main() {
	var ar = []int{4, 4, 6, 7, 8, 9, 2, 3, 1, 19}

	// 快排
	//QuickSort(arr, 9)
	quickSort3Ways(ar, 0, len(ar)-1)
	fmt.Println(ar)
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
