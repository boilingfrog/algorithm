package sort

import "fmt"

var arr = [...]int{2, 1, 99, 6, 88, 3, 23, 545, 80, 223, 89}

//数组指针
var pf *[11]int = &arr

// HeapSort 为了排序使用的是数组指针
func HeapSort() {
	pf[0] = 11
	fmt.Println(pf)
	shiftUp()
}

func shiftUp() {
	fmt.Println(pf)
}
