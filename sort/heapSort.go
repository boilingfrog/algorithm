package sort

import "fmt"

// HeapSort 为了排序使用的是数组指针
func HeapSort(data *[3]int) {
	shiftUp(data)
}

func shiftUp(data *[3]int) {

	fmt.Println(data[2])
}
