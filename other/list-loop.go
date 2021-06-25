package main

/*
单链表中的环是指链表末尾的节点的 next 指针不为 NULL ，而是指向了链表中的某个节点，导致链表中出现了环形结构。
*/

/*
快慢指针
（1）定义两个指针分别为 slow，fast，并且将指针均指向链表头节点。
（2）规定，slow 指针每次前进 1 个节点，fast 指针每次前进两个节点。
（3）当 slow 与 fast 相等，且二者均不为空，则链表存在环。
*/
func isExistLoop(list *ListNode) bool {
	var (
		fast = list
		show = list
	)

	for {
		if show == nil || fast == nil || fast.Next == nil {
			return false
		}
		show = show.Next
		fast = fast.Next.Next
		if fast == show {
			return true
		}
	}
}

/*
哈希缓存法

（1）遍历链表，记录已访问的节点。
（2）将当前节点与之前以及访问过的节点比较，若有相同节点则有环。否则，不存在环。
*/
// TODO

/*
链表反转
*/

func ReverseList(list *ListNode) *ListNode {
	var pre *ListNode
	var cur = list
	for {
		if cur == nil {
			break
		}
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func ReverseListLoop(list *ListNode) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}

	rList := ReverseListLoop(list.Next)
	list.Next.Next = list
	list.Next = nil
	return rList
}
