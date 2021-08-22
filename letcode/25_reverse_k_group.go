package letcode

import "fmt"

// 题目连接 https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

/*
题目描述

给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。

k是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
*/

// 思路
// k个一组进行倒排
// 1、k个一组进行倒叙排，放到一个新的list中
// 2、尾插法将排好的插入存储结果的list中，置空上面倒排的list
// 3、如果不满足k，直接尾插到存储结果的list中

// 主要用到了，链表的头插法和尾插法
// 头插法实现倒排
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	i := 0
	var newList = &ListNode{Val: 0}
	first := ListNode{Val: 0}
	tmp := &first
	now := head

	t := 0
	for {
		if head == nil {
			return first.Next
		}
		if now != nil {
			now = now.Next
			t++
		}

		i++
		if i == k {
			for m := 0; m < k; m++ {
				i--
				if head == nil {
					return first.Next
				}
				fmt.Println(t)
				if t == k {
					node := head
					head = head.Next
					node.Next = newList.Next
					newList.Next = node
				} else {
					tmp.Next = head
					tmp = head
					head = head.Next
				}
			}
			t = 0
			for {
				if newList.Next == nil {
					break
				}
				tmp.Next = newList.Next
				tmp = newList.Next
				newList = newList.Next
			}
		}
	}
}
