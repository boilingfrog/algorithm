package leetcode

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
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := head.Next
	head.Next = swapPairs1(newList.Next)
	newList.Next = head

	return newList
}

// 迭代法

// 本人刚开始的思路就是迭代法
// 不过有一点理解错了，对于node1 := temp.Next中的node1,之后进行的替换，错误的认为是会创建一个新的ListNode进行替换，于是老想着怎么把每次替换的结果放到一起
// 因为是指针，虽然赋值发生了copy,但是是指针指向的地址空间不变，所以进行替换是可以改变原来ListNode中的值
func SwapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
