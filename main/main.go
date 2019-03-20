package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type List struct {
	headNode *ListNode //头节点
}

func main() {

	//res := letcode.Divide(434, 423)
	//fmt.Println(res)

	headNode := &ListNode{}

	listData := headNode
	Insert(1, listData, headNode)
	Insert(2, listData, headNode)
	Insert(4, listData, headNode)

	res := Find(2, listData)
	fmt.Println(res)

	//listData2 := headNode
	//
	//Insert(4, listData2, headNode)
	//Insert(3, listData2, headNode)
	//Insert(2, listData2, headNode)

	//PrintList(listData)

	//res := MergeTwoLists(listData, listData2)
	//fmt.Println(res)
}

//删除节点
func Delete(value int, list *ListNode) {
	pPrev := FindPrevious(value, list)
	_ = pPrev
	p := Find(value, list)
	pPrev = p.Next
	p = nil
}

// ...
func FindPrevious(value int, list *ListNode) *ListNode {
	p := list
	for p.Next != nil && p.Next.Val != value {
		p = p.Next
	}
	return p
}

// 找出链表里面的（不安全）
func Find(value int, list *ListNode) *ListNode {

	p := list
	for p.Val != value {
		p = p.Next
	}
	return p

}

// 测试是否为最后节点 ...
func IsLast(list *ListNode) bool {
	return list.Next == nil
}

// 测试链表是否为空 ...
func isEmpty(list *ListNode) bool {
	return list == nil
}

// 打印链表 ...
func PrintList(list *ListNode) {

	if list.Next != nil {
		fmt.Println(list.Val)
		PrintList(list.Next)
	} else {
		fmt.Println(list.Val)
	}
}

// 合并两个有序的单链表 ...
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l2
		res.Next = MergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = MergeTwoLists(l1.Next, l2)
	}
	return res
}

// 插入节点 头插法
func Insert(value int, list *ListNode, position *ListNode) {
	tempCell := new(ListNode)
	//fmt.Println("++++", tempCell)
	if tempCell == nil {
		fmt.Println("out of space!")
	}
	tempCell.Val = value
	tempCell.Next = position.Next
	position.Next = tempCell
}
