package main

func bsearch1(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// ---------------------------------------------------------------------------------------------------------------------
// 递归的实现

func bsearchDigui(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return bsearchDigui(arr, target, left, right)
}

func bsearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	return bsearchDigui(arr, target, 0, len(arr)-1)
}

// ---------------------------------------------------------------------------------------------------------------------
// 查找第一个值等于给定值的元素
func findFirstEqual(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		if mid >= 0 && arr[mid-1] == target {
			right = mid - 1
		} else {
			return mid
		}
	} else if arr[mid] < target {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return findFirstEqual(arr, target, left, right)
}

func bsearchFirstEqual(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	return findFirstEqual(arr, target, 0, len(arr)-1)
}

// ---------------------------------------------------------------------------------------------------------------------
// 查找最后一个值等于给定值的元素
func findLastEqual(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		if mid < len(arr) && arr[mid+1] == target {
			left = mid + 1
		} else {
			return mid
		}
	} else if arr[mid] < target {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return findLastEqual(arr, target, left, right)
}

func bsearchLastEqual(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	return findLastEqual(arr, target, 0, len(arr)-1)
}
