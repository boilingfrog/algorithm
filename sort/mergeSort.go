package sort

// 归并排序
func MergeSort(arr []int, n int) {

	mergeSort(arr, 0, n-1)
}

// 递归使用归并排序，对arr[1.......r]的范围进行排序
func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2 //问题点可能这个地方会出现内存的溢出
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	if arr[mid] > arr[mid+1] {
		Merge(arr, l, mid, r)
	}
}

