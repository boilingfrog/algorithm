package leetcode

import "sort"

/*
给定一个区间的集合intervals，其中 intervals[i] = [starti, endi]。返回 需要移除区间的最小数量，使剩余区间互不重叠。


示例 1:

输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。


示例 2:

输入: intervals = [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。



示例 3:

输入: intervals = [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-overlapping-intervals
*/

/*
思路

贪心算法

移除最小的区间，保证无重叠的区域最大

根据区域的右节点进行从小到大的排序

循环遍历，取第一个节点作为当前节点

循环的下一个节点，左节点大于等于，当前节点的值，说明无重合

*/
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	maxNum, interval := 1, intervals[0][1]
	for _, item := range intervals[1:] {
		if item[0] >= interval {
			interval = item[1]
			maxNum++
		}

	}

	return n - maxNum
}
