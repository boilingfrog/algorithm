package main

import "fmt"

func QuickSort(arr []int, n int) {

	quickSort(arr, 0, n-1)

	fmt.Println(arr)

}

// 对arr[l.......r]部分进行快速排序
func quickSort(arr []int, l int, r int) {

	if l >= r {
		return
	}

	p := partition(arr, l, r)
	fmt.Println(p)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

//返回p使得arr[l...p-1]<arr[p];arr[p+1....r]>arr[p]
func partition(arr []int, l int, r int) int {
	v := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			Swap(arr, j+1, i)
			t := arr[j+1]
			arr[j+1] = arr[i]
			arr[i] = t
			j++
		}
	}
	t := arr[l]
	arr[l] = arr[j]
	arr[j] = t
	return j + 1
}
