package main

import (
	"suanfa/sort"
)

func main() {
	testSlice := []int{1, 2, 3, 42, 3, 1, 4, 5, 6, 2, 1, 23, 43, 1, 32, 3, 5, 2, 1, 8, 54, 4, 0}
	sort.HeapSort(testSlice)
}
