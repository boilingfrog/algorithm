package main

import "fmt"

//三路快速排序处理，arr[l.....r]
//将arr[l...r]分三部分<v.=V >V
// 之后将<v.>v两部分进行三路快速排序

/*
笨如狗

这个小快排。竟然写了好久才完成，害

总结了递归的写法

- 1.一个问题的解可以分解为几个子问题的解

- 2.这个问题与分解之后的子问题，除了数据规模不同，求解思路完全一样

- 3.存在递归终止条件

就是通用方法抽离，然后找到终止的条件
*/
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
