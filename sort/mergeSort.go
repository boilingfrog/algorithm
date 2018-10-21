package sort

import "fmt"

// 归并排序   
func MergeSort(arr []int, n int) {

	mergeSort(arr, 0, n-1)

	fmt.Println(arr)

}

// 递归使用归并排序，对arr[1.......r]的范围进行排序
func mergeSort(arr []int, l int, r int) {

	if l >= r {
		return
	}
	mid := (l + r) / 2 //问题点可能这个地方会出现内存的溢出
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

// 将arr[l......mid]和arr[mid+1.....r]两部分进行归并
func merge(arr []int, l int, mid int, r int) {
	num := r - l + 1
	aux := make([]int, num)


	//var aux [r - l + 1]int //声明临时的空间

	for i := l; i <= r; i++ {
		aux[i-l] = arr[i] //l到r之间，放到临时的数组空间，0到l之间有一个l的偏移量
	}

	i := l
	j := mid + 1

	for k := l; k <= r; k++ {

		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}

	return
}
