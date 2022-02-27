package leetcode

/*
给你一个非负整数数组nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。


示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
    从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。


示例 2:

输入: nums = [2,3,0,1,4]
输出: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-ii
*/

/*
思路

在每次可跳动的范围内找出，能跳动的最大值，然后到了最大可跳动的值，增加跳动的步骤
*/
func jump(nums []int) int {
	steps := 0
	end := 0
	maxTarget := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxTarget {
			maxTarget = i + nums[i]
		}

		// 这一步是精华
		// 到了上一步跳转能到的最大范围值的时候，增加跳转的步骤
		if end == i {
			end = maxTarget
			steps++
		}
	}

	return steps
}

/*
是个笨办法，两层循环

时间复杂度是O(n^2)

从最后的一个节点开始循环，然后找到能跳转到当前节点的最小索引的节点，然后更新当前节点到刚刚找到的最小节点，依次循环。。。。。
*/
func jump1(nums []int) int {
	steps := 0
	position := len(nums) - 1

	for position > 0 {
		for i := 0; i < position; i++ {
			if nums[i]+i >= position {
				position = i
				steps++
				break
			}
		}
	}

	return steps
}
