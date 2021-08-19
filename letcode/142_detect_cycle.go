package letcode

import "fmt"

/*
地址：https://leetcode-cn.com/problems/linked-list-cycle-ii/


给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：

你是否可以使用 O(1) 空间解决此题？


输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
*/

// 具体的题解参考https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head
	quick := head.Next.Next

	for {
		if quick == nil || quick.Next == nil {
			return nil
		}
		if quick == slow {
			slow1 := head
			for slow1 != slow {
				slow1 = slow1.Next
				slow = slow.Next
				fmt.Println(slow1, slow)
			}
			return slow1
		}
		slow = slow.Next
		quick = quick.Next
		if quick != nil && quick.Next != nil {
			quick = quick.Next
		}
	}
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 便利报错的map中，如果map中有重复，返回即可
func detectCycle1(head *ListNode) *ListNode {
	list := head
	var cacheMap = make(map[*ListNode]int)
	for {
		if list == nil || list.Next == nil {
			return nil
		}
		if _, ok := cacheMap[list]; ok {
			return list
		}
		cacheMap[list] = list.Val
		list = list.Next
	}
}
