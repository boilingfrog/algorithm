package main

import "fmt"

func main() {
	fmt.Println(212)

	// list

	headNode := &ListNode{}

	listData := headNode
	Insert(1, listData, headNode)
	Insert(2, listData, headNode)
	Insert(4, listData, headNode)

	PrintList(listData)

	fmt.Println(DoublePointer(listData, 2))
}
