package letcode

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。



示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。


进阶：

如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/
func minSubArrayLen(s int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var (
		num   = 0
		left  = 0
		right = 0
		sum   = 0
	)

	for left < len(nums) {

		if right < len(nums) && sum < s {
			sum += nums[right]
			right++
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= s {
			if num == 0 {
				num = right - left
			} else {
				if num > (right - left) {
					num = right - left
				}

			}
		}
	}

	return num
}
