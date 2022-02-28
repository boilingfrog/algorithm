package leetcode

/*
316. 去除重复字母
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。



示例 1：

输入：s = "bcabc"
输出："abc"
示例 2：

输入：s = "cbacdcbc"
输出："acdb"

题目地址 https://leetcode-cn.com/problems/remove-duplicate-letters/
*/

/*
思路

这里使用了栈，每次去处理一个字符串，先去栈中找出上一个，如果栈中的序列小于当前的字符串，并且栈中的在后面还有重复的字符串，弹出栈中元素，后面乜有重复的，当前元素入栈，依次循环

*/

func removeDuplicateLetters(s string) string {
	stack := []rune{}
	isStack := make(map[rune]int, 0)
	isInStack := make(map[rune]bool, 0)

	for _, item := range s {
		isStack[item]++
	}

	for _, item := range s {
		for len(stack) > 0 && stack[len(stack)-1] > item && isStack[stack[len(stack)-1]] > 1 && !isInStack[item] {
			isStack[stack[len(stack)-1]]--
			isInStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}

		if !isInStack[item] {
			stack = append(stack, item)
			isInStack[item] = true
		} else {
			isStack[item]--
		}
	}

	return string(stack)
}
