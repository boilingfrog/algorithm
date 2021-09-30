package main

import "fmt"

//三路快速排序处理，arr[l.....r]
//将arr[l...r]分三部分<v.=V >V
// 之后将<v.>v两部分进行三路快速排序
func QuickSort3Ways(arr []int, n int) {

	quickSort3Ways(arr, 0, n-1)

	fmt.Println(arr)

}

func quickSort3Ways(arr []int, lo, hi int) {
	if len(arr) < 2 || lo == hi {
		return
	}

	v := arr[lo]
	lt, gt := lo, hi
	i := lo + 1
	for lt < gt {
		if v == arr[i] {
			i++
		} else if v > arr[i] {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
			i++
		} else if v < arr[i] {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		}
	}
	arr[gt+1], arr[i] = arr[i], arr[gt+1]

	quickSort3Ways(arr, lo, lt)

	quickSort3Ways(arr, gt+1, hi)
}
