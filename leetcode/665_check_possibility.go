package leetcode

import "sort"

/*
给你一个长度为n的整数数组nums，请你判断在 最多 改变1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的：对于数组中任意的i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

示例 1:

输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个 4 变成 1 来使得它成为一个非递减数列。


示例 2:

输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
*/

/*
思路

这里借用了 sort 包中的函数

假定修改一次就可以了，所以可以修改 i 或 i+1

然后判断修改之后的数组是否是已经排好序的了

时间复杂度 O(N)

空间复杂度 O(1)
*/

func checkPossibility(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}

	for i := 0; i < n-1; i++ {
		lef, rig := nums[i], nums[i+1]
		if lef > rig {
			nums[i] = rig
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = lef
			nums[i+1] = lef

			return sort.IntsAreSorted(nums)
		}
	}

	return true
}
