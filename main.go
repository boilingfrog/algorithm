package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 3, 2, 3, 4, 6, 23, 34, 56, 78, 102}
	fmt.Println(MergeSort(nums))
}

func MergeSort(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	var tmp []int
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			tmp = append(tmp, nums[j])
			j++
		} else {
			tmp = append(tmp, nums[i])
			i++
		}
	}
	if i <= mid {
		tmp = append(tmp, nums[i:mid+1]...)
	}
	if j <= end {
		tmp = append(tmp, nums[j:end+1]...)
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
}

//func mergeSort(arr []int, l int, r int) int {
//	if l >= r {
//		return 0
//	}
//
//	mid := l + (r-l)/2
//	cnt := mergeSort(arr, l, mid)
//	cnt += mergeSort(arr, mid+1, r)
//	//if arr[mid] > arr[mid+1] {
//	cnt += Merge(arr, l, mid, r)
//	//}
//
//	return cnt
//}
//
//func Merge(arr []int, l int, mid int, r int) int {
//	num := r - l + 1
//	aux := make([]int, num)
//	for i := l; i <= r; i++ {
//		aux[i-l] = arr[i] //l到r之间，放到临时的数组空间，0到l之间有一个l的偏移量
//	}
//	i := l
//	j := mid + 1
//	cnt := 0
//	for k := l; k <= r; k++ {
//		if i > mid {
//			arr[k] = aux[j-l]
//			j++
//			cnt++
//		} else if j > r {
//			arr[k] = aux[i-l]
//			i++
//		} else if aux[i-l] < aux[j-l] {
//			arr[k] = aux[i-l]
//			i++
//		} else {
//			arr[k] = aux[j-l]
//			j++
//			cnt++
//		}
//	}
//	return cnt
//}

func test(n int) [][]int {
	var sil = make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sil[i] = append(sil[i], j)
		}
	}
	return sil
}

func sum(n, m int) int {
	sum := 0
	for i := 0; i < n; i++ {
		if m == n {
			return sum
		}
		sum += i * 2
	}
	return sum
}
