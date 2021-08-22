package main

func main() {
	// list

	headNode := &ListNode{}

	listData := headNode
	Insert(1, listData, headNode)
	Insert(2, listData, headNode)

	// 尾插法
	head := ListNode{}
	temp := &head
	for i := 1; i < 10; i++ {
		stu := ListNode{Val: i}
		temp.Next = &stu
		temp = &stu
	}

	PrintList(&head)
	//
	//fmt.Println(DoublePointer(listData, 2))
	//
	//fmt.Println(isExistLoop(listData))

}

func appendNode(first, list *ListNode) *ListNode {
	if list.Next != nil {
		first.Next = list.Next
		list = list.Next
		appendNode(first.Next, list)
	}

	return first
}
