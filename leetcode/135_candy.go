package leetcode

/*
135. 分发糖果

n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



示例1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。



示例2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/candy
*/

/*
思路

两遍循环，先左循环一遍，找出右节点和左节点的关系，如果右节点大于左节点left[i-1] + 1，计算出一遍值

然后从右到左循环一遍，找出左节点和右节点的关系，如果左节点大于右节点target = target + 1，计算出一遍值

然后循环比较两次对比的值，找出最大的值进行累加

时间负载度 O(N)

空间复杂度 O(N)
*/

func candy(ratings []int) int {
	nums := len(ratings) - 1
	candies := 0
	left := make([]int, nums+1)

	for i := 0; i <= nums; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	target := 0
	for i := nums; i >= 0; i-- {
		if i < nums && ratings[i] > ratings[i+1] {
			target = target + 1
		} else {
			target = 1
		}

		if left[i] > target {
			candies += left[i]
		} else {
			candies += target
		}
	}

	return candies
}
