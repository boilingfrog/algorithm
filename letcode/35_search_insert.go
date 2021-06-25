package letcode

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

解题思路

1、暴力法

一层循，设置边界值

2、二分法


*/

// 暴力法
func searchInsert1(nums []int, target int) int {
	var flg bool
	for index, item := range nums {
		if item < target {
			flg = true
		}

		if item >= target && index > 0 && flg {
			return index
		}

		if index == len(nums)-1 && flg == true {
			return index + 1
		}
	}

	return 0
}

// 二分查找法
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := right + (left-right)/2
		fmt.Println(left, right, mid)
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}

	}

	return left
}
