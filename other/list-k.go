package main

/*
题目：输入一个单链表，输出此链表中的倒数第 K 个节点。（去除头结点，节点计数从 1 开始）。
*/

/*
思路：

1、两次遍历，找到链表的长度，然后总数-k，就是正序的节点，然后遍历找到这个几点的信息

2、递归
（1）定义num = k
（2）使用递归方式遍历至链表末尾。
（3）由末尾开始返回，每返回一次 num 减 1
（4）当 num 为 0 时，即可找到目标节点

3、双指针法
（1）定义两个指针 p1 和 p2 分别指向链表头节点。
（2）p1 前进 K 个节点，则 p1 与 p2 相距 K 个节点。
（3）p1，p2 同时前进，每次前进 1 个节点。
（4）当 p1 指向到达链表末尾，由于 p1 与 p2 相距 K 个节点，则 p2 指向目标节点。

*/

// 两次遍历
func Len(list *ListNode) int {
	var count int

	for list.Next != nil {
		list = list.Next
		count++
	}

	return count
}

// 递归(有问题)
func FindRecursion(list *ListNode, k int) *ListNode {
	num := k
	_ = num
	if list == nil {
		return nil
	}
	listNow := FindRecursion(list.Next, k)

	if listNow != nil {
		return listNow
	} else {
		num--
		if num == 0 {
			return list
		}
	}
	return list
}

// 双指针法
func DoublePointer(list *ListNode, k int) *ListNode {
	if list == nil || k == 0 {
		return nil
	}

	var (
		p1 = list
		p2 = list
	)

	// p1先走k个节点
	for i := 0; i < k; i++ {
		if p1 != nil {
			p1 = p1.Next
		} else {
			return nil
		}
	}

	for {
		if p1 != nil {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			return p2
		}
	}

	return p2
}
