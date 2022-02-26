package leetcode

/*

给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

示例1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。


示例2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
*/

/*
时间复杂度 O(N)

空间复杂度 O(1)
*/

/*
解题思路：刚开始的考虑，想着按照题目的思路。列出所有的可能性，然后一个个对应起来思考解题的方法

但是这种方法，每次做算法题目好像都不太行，应为有太多的可能性需要处理了

看了答案，很奇妙，处理思路，没有盯着一个个的可能的case去处理，而是把问题的处理方案给是升华了，抓住了问题的根本

思考 要善于总结，面对这种题目不能纠结于每一个case，应该善于总结case的共通点，进而找到破解的思路
*/

/*
思路一

能跳转到最后，就意味着，遍历所有的节点，然后计算出当前节点的能跳转的最大长度，如果这个长度小于当前节点的下标，就表示不能跳转到当前节点了，也就是不能跳转到最后的节点了
*/

func canJump(nums []int) bool {
	target, num := 0, len(nums)
	for i := 0; i < num-1; i++ {
		if i > target {
			return false
		}

		if nums[i]+i > target {
			target = nums[i] + i
		}
	}

	if target >= num-1 {
		return true
	}

	return false
}

/*
思路二

反向从最后一个节点开始向前走

从最后一个节点开始循环，记录下标节点的为目标节点，如果前一个节点的值加上前一个节点的下标索引值能够大于当前节点的下标值，说明前一个节点可以跳转到当前节点，更新目标节点，然后依次对比


如果循环一遍，目标节点的值>0，说明调不到最后的节点
*/

func canJump1(nums []int) bool {
	target, num := len(nums)-1, len(nums)-1
	for i := num - 1; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}

	if target <= 0 {
		return true
	}

	return false
}
