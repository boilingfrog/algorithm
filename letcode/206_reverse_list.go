package letcode

// 题目地址：https://leetcode-cn.com/problems/reverse-linked-list/

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]


示例 2：

输入：head = [1,2]
输出：[2,1]


示例 3：

输入：head = []
输出：[]

*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newList = &ListNode{Val: 0}
	for {
		if head == nil {
			return newList.Next
		}
		node := head
		head = head.Next
		node.Next = newList.Next
		newList.Next = node
	}
}
