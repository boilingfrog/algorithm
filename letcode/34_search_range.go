package letcode

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func SearchRange(nums []int, target int) []int {
	fmt.Println(nums)
	var res = []int{-1, -1}

	resData := pointsDigui(nums, target, 0, len(nums)-1, res)
	return resData
}

func pointsDigui(nums []int, target int, low int, high int, res []int) []int {

	min := (low + high) / 2

	if nums[min] == target && nums[min-1] < target {
		res[0] = min
		low = min + 1
	} else if nums[min] == target && nums[min+1] > target {
		res[1] = min
		high = min - 1
	} else if target > nums[min] {
		low = min + 1
	} else if target < nums[min] {
		high = min - 1
	}

	if low == high {
		return res
	} else {
		fmt.Println(low)
		fmt.Println(high)

		return pointsDigui(nums, target, low, high, res)
	}
}
