package main

import (
	"fmt"
	"suanfa/letcode"
)

func main() {
	list1 := letcode.ListNode{}
	list1.ListAdd(1)

	//travselLinkList(&list1)

	list2 := letcode.ListNode{}
	list2.ListAdd(9)
	list2.ListAdd(9)

	//travselLinkList(&list2)

	res := letcode.AddTwoNumbers(&list1, &list2)
	_ = res

	//fmt.Println(10 % 10)
	//fmt.Println(6 / 10)

	travselLinkList(res)
}

func travselLinkList(list *letcode.ListNode) {
	//遍历
	head := list
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
