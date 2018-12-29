package letcode

import "fmt"

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
func LengthOfLongestSubstringSilnce(s string) int {

	mapp := make(map[string]int)
	mappString := make(map[int]string)
	index := 1
	for i := 0; i < len(s); i++ {

		if _, ok := mapp[s[i:i+1]]; ok {
			index++

			mappString[index] = s[mapp[s[i:i+1]]+1 : i+1]
			if mapp[s[i:i+1]]+1 == i {
				mapp = make(map[string]int)

			}

			mapp[s[i:i+1]] = i

		} else {
			mapp[s[i:i+1]] = i
			mappString[index] += string(s[i : i+1])
		}
		fmt.Println(mappString)
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
