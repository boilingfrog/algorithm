package main

import "suanfa/sort"

func main() {
	var arr = [...]int{2, 1, 999}
	//数组指针
	var pf *[3]int = &arr
	_ = pf
	sort.HeapSort(pf)
}
