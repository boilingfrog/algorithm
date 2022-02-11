package leetcode

import "sort"

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。



示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。


进阶：

如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

思路

过了几个月重新做了一次还是不会，害，尴尬

总结下思路

刚开始的思路也想到了双指针，一个左节点一个又节点一起移动，想的是右边节点移动，如果右边节点遇到比当前遇到的值，都要大的值之后，然后左边节点移动，是个错误的想法。。。。。


答案很好，如果和没有达到目标值，右边节点移动，如果达到了，记录下当前的两数间隔，然后移动左节点，依次循环。。。  实在是秒啊
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

func minSubArrayLenBinary(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 计算当前节点之前所有节点的和
	// 这里使用了定义n+1,也就是在计算节点之前放置一个0，方便原来第一个节点的比较
	numsTarget := make([]int, n+1)
	for i := 1; i <= n; i++ {
		numsTarget[i] = numsTarget[i-1] + nums[i-1]
	}

	res := 0
	for i := 1; i < n+1; i++ {
		// 找出大于当前节点+target的数，在计算数组numsTarget出现的第一个位置，这个位置减去当前节点的下标就是，我们需要的范围值
		targetData := numsTarget[i-1] + target
		bound := sort.SearchInts(numsTarget, targetData)

		if bound == 0 || bound == len(numsTarget) {
			continue
		}
		// 找出最小的
		if bound > 0 {
			bound = bound - (i - 1)
			if res == 0 {
				res = bound
			} else if bound < res {
				res = bound
			}
		}
	}

	return res
}
