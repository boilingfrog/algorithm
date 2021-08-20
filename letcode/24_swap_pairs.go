package letcode

/*
题目地址：https://leetcode-cn.com/problems/swap-nodes-in-pairs/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解析思路：传统的递归
// 刚开始的想法向通过for循环来处理，但是试了好久没成功，每次循环的，总会产生新的指针，不能在一个指针中进行操作

// 递归的思路很好，每次处理两个，预留改组的next是下一个递归的返回，知道所有的都成功
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := head.Next
	head.Next = swapPairs(newList.Next)
	newList.Next = head

	return newList
}
