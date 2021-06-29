package letcode

import "sort"

/*
给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。


示例 2：

输入：nums = [1]
输出：1


示例 3：

输入：nums = [0]
输出：0


示例 4：

输入：nums = [-1]
输出：-1


示例 5：

输入：nums = [-100000]
输出：-100000


提示：

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105


进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/

// 动态规划
// 若前一个元素大于0，则将其加到当前元素上
// 计算中所有的可能值，然后找出最大的即可
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var targets = make([]int, len(nums))
	targets[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if targets[i-1] < 0 {
			targets[i] = nums[i]
		} else {
			targets[i] = targets[i-1] + nums[i]
		}
	}
	sort.Ints(targets)
	return targets[len(nums)-1]
}

// 贪心
// 若当前指针所指元素的和小于0，则丢弃当前元素之前的数列
func MaxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var target, now = nums[0], 0
	for i := 0; i < len(nums); i++ {
		now = Max1(nums[i], nums[i]+now)
		target = Max1(now, target)
	}

	return target
}

func Max1(i, j int) int {
	if i < j {
		return j
	}
	return i
}
