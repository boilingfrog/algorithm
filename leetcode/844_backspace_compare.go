package leetcode

import (
	"fmt"
	"strings"
)

/*
844. 比较含退格的字符串
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。



示例 1：

输入：S = "ab#c", T = "ad#c"
输出：true
解释：S 和 T 都会变成 “ac”。


示例 2：

输入：S = "ab##", T = "c#d#"
输出：true
解释：S 和 T 都会变成 “”。


示例 3：

输入：S = "a##c", T = "#a#c"
输出：true
解释：S 和 T 都会变成 “c”。


示例 4：

输入：S = "a#c", T = "b"
输出：false
解释：S 会变成 “c”，但 T 仍然是 “b”。


提示：

1 <= S.length <= 200
1 <= T.length <= 200
S 和 T 只含有小写字母以及字符 '#'。


进阶：

你可以用 O(N) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？
*/

// 思路1使用栈进行啊操作，时间复杂度和空间复杂度都是O(N+M)
func backspaceCompare1(s string, t string) bool {
	if s == t {
		return true
	}

	stack1 := Init(10)
	stack2 := Init(10)
	dealData(s, stack1)
	dealData(t, stack2)

	for !stack1.IsEmpty() || !stack2.IsEmpty() {
		fmt.Println(stack1.Pop(), stack2.Pop())
		if stack1.Pop() != stack2.Pop() {
			return false
		}
	}

	if stack1.IsEmpty() && stack2.IsEmpty() {
		return true
	}
	return false

}

func dealData(d string, stack *stack) {
	for _, item := range strings.Split(d, "") {
		if item == "#" {
			if !stack.IsEmpty() {
				stack.Pop()
			}
		} else {
			stack.Push(item)
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// 双指针
// 反思，没有想出来，整体的思路是有了，边界值没有处理好，
// 对于#的回退没有处理好，原来想法是遇到一个#，然后回退的count++，但是仔细推敲了，这种针对aa#a##,这种场景就不能满足了
// 看了官方的题解 引入了 skipS，遇到# ，就skipS++，最后再去处理skipS，这个用的秒，值得学习
func backspaceCompare(s string, t string) bool {
	if s == t {
		return true
	}

	i := len(s) - 1
	j := len(t) - 1
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}

	return true
}
