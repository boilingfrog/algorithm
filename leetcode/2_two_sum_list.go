package leetcode

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resNode := ListNode{}

	head1 := l1.Next
	head2 := l2.Next
	resX := 0

	for head1 != nil || head2 != nil {
		res := 0
		if head1 != nil {
			res += head1.Val
			head1 = head1.Next
		}

		if head2 != nil {
			res += head2.Val
			head2 = head2.Next
		}
		resY := res

		if res >= 10 {
			resY = resY - 10
		}
		resZ := resY + resX

		//resNode.ListAdd(resY + resX)
		if resZ >= 10 {
			resZ = resZ - 10
			res = resZ
		}
		fmt.Println(resZ)
		fmt.Println("++++++++++++++++++++++")

		node := &ListNode{Val: resZ}
		//
		//node.Next = resNode.Next
		//resNode.Next = node

		if resNode.Next == nil {
			resNode.Next = node
		} else {
			cur := resNode.Next
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Next = node
		}

		resX = res / 10

	}
	for resX > 0 {
		resY := resX
		if resX >= 10 {
			resY = resY - 10
		}

		node := &ListNode{Val: +resX}
		//
		//node.Next = resNode.Next
		//resNode.Next = node

		if resNode.Next == nil {
			resNode.Next = node
		} else {
			cur := resNode.Next
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Next = node
		}
		resX = resX / 10

	}

	return &resNode
}

//从头部添加元素
func (this *ListNode) ListAdd(data int) {
	node := &ListNode{Val: data}
	node.Next = this.Next
	this.Next = node
}

// 单链表的长度
func (this *ListNode) ListLength() int {
	cur := this.Next
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}
