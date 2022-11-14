package main

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := MergeSort(arr[0:mid])
	rightArr := MergeSort(arr[mid:])
	return merge(leftArr, rightArr)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	var result []int
	// 合并子序列
	for {
		// 任何一个区间遍历完成就退出
		if i >= m || j >= n {
			break
		}
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}

	// 左边的子集没有遍历完成，将左侧的放入到结果集
	if i < m {
		result = append(result, left[i:]...)
	}

	// 右侧的子集没有遍历完成，将右侧的放入到结果集中
	if j < n {
		result = append(result, right[j:]...)
	}

	return result
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
