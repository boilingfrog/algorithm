package main

/*合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	sorted := &ListNode{Val: 0, Next: nil}
	tailNum := 0
	minKey := 0
	cur := make([]*ListNode, len(lists))
	for i := 0; i < len(lists); i++ {
		cur[i] = lists[i]
	}
	for tailNum < len(lists) {
		var min int
		minFlag := false
		for i := 0; i < len(lists); i++ {
			if cur[i] == nil {
				continue
			}
			if minFlag == false {
				min = cur[i].Val
				minKey = i
				minFlag = true
				continue
			}

			l := cur[i]
			if l.Val < min {
				min = l.Val
				minKey = i

			}
		}
		if minFlag == false {
			break
		}
		cur[minKey] = cur[minKey].Next
		if cur[minKey] == nil {
			tailNum++
		}
		if head.Next == nil {
			sorted = &ListNode{Val: min, Next: nil}
			head.Next = sorted
		} else {
			sorted.Next = &ListNode{Val: min, Next: nil}
			sorted = sorted.Next

		}

	}
	return head.Next
}
