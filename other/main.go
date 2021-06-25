package main

func main() {
	// list

	headNode := &ListNode{}

	listData := headNode
	Insert(1, listData, headNode)
	Insert(2, listData, headNode)
	Insert(4, listData, headNode)

	//PrintList(listData)
	//
	//fmt.Println(DoublePointer(listData, 2))
	//
	//fmt.Println(isExistLoop(listData))

	PrintList(listData)

	// PrintList(ReverseList(listData))
	PrintList(ReverseListLoop(listData))

}
