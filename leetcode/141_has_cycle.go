package leetcode

/*
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。


*/
// 很low的一种方法，将其放到map中，循环遍历，如果在map中能再次找到这个，说明有值了
func hasCycle1(head *ListNode) bool {
	var cacheMap = make(map[*ListNode]int)
	for {
		if head == nil || head.Next == nil {
			return false
		}
		if _, ok := cacheMap[head]; ok {
			return true
		}
		cacheMap[head] = head.Val
		head = head.Next
	}
}

// 使用快慢指针来处理下
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	quick := head.Next

	for {
		if quick == nil || head.Next == nil || slow == nil || slow.Next == nil {
			return false
		}

		if quick == slow {
			return true
		}
		slow = slow.Next
		quick = quick.Next
		if quick != nil && quick.Next != nil {
			quick = quick.Next
		}
	}
}
