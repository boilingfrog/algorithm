package sort

import "fmt"

//三路快速排序处理，arr[l.....r]
//将arr[l...r]分三部分<v.=V >V
// 之后将<v.>v两部分进行三路快速排序
func QuickSort3Ways(arr []int, n int) {

	quickSort3Ways(arr, 0, n-1)

	fmt.Println(arr)

}

// 对arr[l.......r]部分进行快速排序
func quickSort3Ways(arr []int, l int, r int) {
	v := arr[l]
	lt := l
	gt := r + 1
	i := l + 1
	if i > gt {
		return
	}
	// 随机生成一个数作为标定点，定为3
	Swap(arr, l, 5)


	for {
		if i >= gt {
			break
		}
		fmt.Println(i-gt)
		if arr[i] < v {
			Swap(arr, l, i+1)
			lt++
			i++
		} else if arr[i] > v {
			Swap(arr, l, gt-1)
			gt--
		} else {
			i++
		}
	}

	Swap(arr, l, lt)
	quickSort3Ways(arr,l,lt-1)
	quickSort3Ways(arr,gt,r)

}