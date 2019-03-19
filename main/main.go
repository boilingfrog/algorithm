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
	list := ListNode{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	fmt.Println(list)

	//res := letcode.MergeTwoLists(list., list1)
	//fmt.Println(res)
}

//从头部添加元素
func (this *ListNode) Add(data int) {
	node := &ListNode{Val: data, Next: this.Next}
	this.Val = data

	this.Next = node
}

//插入节点
func Insert(value int, list **ListNode) {
	stu := ListNode{
		Val: 2,
	}

	stu.Next = *list

	*list = &stu

}
