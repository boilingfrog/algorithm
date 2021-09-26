package letcode

import (
	"strings"
)

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true


示例 2：
输入：s = "()[]{}"
输出：true

示例3：
输入：s = "(]"
输出：false


示例4：
输入：s = "([)]"
输出：false


思路：


刚开始对于题目的理解有点错误了，认为题目的中的括号会嵌套出现，结果不是，没认证阅读题目，导致自己想的复杂了。

既然是没有嵌套就简单多了，使用栈就是可以了，利用后进先出，能够实现一对对的匹配。
*/

func isValid(s string) bool {
	stack := Init(10)
	for _, item := range strings.Split(s, "") {
		switch item {
		case "{":
			stack.Push("{")
		case "(":
			stack.Push("(")
		case "[":
			stack.Push("[")
		case "}":
			if stack.Pop() != "{" {
				return false
			}
		case ")":
			if stack.Pop() != "(" {
				return false
			}
		case "]":
			if stack.Pop() != "[" {
				return false
			}
		}
	}

	if stack.Pop() == "-1" {
		return true
	}

	return false
}

type stack struct {
	Sil   []string
	Count int
	N     int
}

func Init(n int) *stack {
	return &stack{
		N:   n,
		Sil: make([]string, n),
	}
}

func (s *stack) Push(data string) {
	if len(s.Sil) < (s.Count + 1) {
		s.Sil = append(s.Sil, data)
	} else {
		s.Sil[s.Count] = data

	}
	s.Count++
}

func (s *stack) Pop() string {
	if s.Count == 0 {
		return "-1"
	}
	s.Count--
	return s.Sil[s.Count]
}
