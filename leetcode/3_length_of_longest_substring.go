package leetcode

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

*/

//  pwwkew
// ""wobgrovw""
// abcabcbb

/*

作者：Mono_Chrome
链接：https://www.nowcoder.com/discuss/357566?type=1
来源：牛客网

题目：给定m个不重复的字符 [a, b, c, d]，以及一个长度为n的字符串tbcacbdata，问能否在这个字符串中找到一个长度为m的连续子串，使得这个子串刚
好由上面m个字符组成，顺序无所谓，返回任意满足条件的一个子串的起始位置，未找到返回-1。比如上面这个例子，acbd，3。

*/
func lengthOfLongestSubstring2(s string) int {

	return 1
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

/*
 使用滑动窗口
*/
func lengthOfLongestSubstring1(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}

func LengthOfLongestSubstringGufang(s string) int {
	if s == "" {
		return 0
	}
	i, j, max := 0, 1, 0
	for j <= len(s) {
		if j == len(s) {
			max = Max(max, j-i)
			break
		}
		x := strings.IndexByte(s[i:j], s[j])
		max = Max(max, j-i)
		if x != -1 {
			i = i + x + 1
		}
		j++
	}
	return max
}

func LengthOfLongestSubstringSilnce(s string) int {

	mapp := make(map[string]int)
	mappString := make(map[int]string)
	index := 1
	for i := 0; i < len(s); i++ {

		if _, ok := mapp[s[i:i+1]]; ok {
			index++
			mappString[index] = s[mapp[s[i:i+1]]+1 : i+1]
			mapp = make(map[string]int)

			mapp[s[i:i+1]] = i

		} else {
			mapp[s[i:i+1]] = i
			mappString[index] += string(s[i : i+1])
		}
		fmt.Println(mapp)
	}
	lenNum := 0
	for _, item := range mappString {

		if len(item) > lenNum {
			lenNum = len(item)
		}
	}
	return lenNum
}

func LengthOfLongestSubstring(s string) int {

	mapp := make(map[int32]int32)
	silence := make(map[int]string)
	index := 1
	var str int32
	_ = str
	for key, item := range s {

		if _, ok := mapp[item]; ok {
			index++

			if key > 0 {
				if item != str {
					silence[index] += string(str)
				}
				silence[index] += string(item)
			}
			mapp := make(map[int32]int32)
			_ = mapp
		} else {
			mapp[item] = item
			silence[index] += string(item)
			str = item
		}
	}
	lenNum := 0
	for _, item := range silence {

		if len(item) > lenNum {
			lenNum = len(item)
		}
	}

	return lenNum
}
