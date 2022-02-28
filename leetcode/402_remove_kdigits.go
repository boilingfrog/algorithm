package leetcode

import "strings"

/*
402. 移掉 K 位数字
给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。


示例 1 ：

输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。


示例 2 ：

输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。


示例 3 ：

输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。


题目地址：https://leetcode-cn.com/problems/remove-k-digits/
*/

/*
思路

这种题目是有规律的

从左往右开始遍历，如果i-1 对应的值大于i对应的值，直接删除i-1的值即可，如果不存在，说明整个数字序列单调不降，删去最后一个数字即可。

对于值的删除过程使用栈来处理，先入栈一个值，然后从num拿出一个元素，如果小于栈末元素，弹出改值，然后对比的当前值入栈，依次循环

时间复杂度 O(n)

空间复杂度 O(n)
*/

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for _, item := range num {
		for k > 0 && len(stack) > 0 && byte(item) < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, byte(item))
	}

	stack = stack[:len(stack)-k]

	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
