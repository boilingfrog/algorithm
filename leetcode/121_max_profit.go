package leetcode

/*
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。


示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
*/

/*
思路

刚开始想到了蓄水池的问题，仔细的想了想发现指针的移动是个问题

遍历的过程中寻找一个最小的买进价格，然后每次和最小的买进价格进行计算，找出最大的收益值
*/

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	target := prices[0]
	price := 0

	for _, item := range prices[1:] {
		if target > item {
			target = item
		}

		if price < (item - target) {
			price = item - target
		}

	}

	return price
}
