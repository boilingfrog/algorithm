package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type ListNode1 struct {
	Val  int
	Next *ListNode1
}
type List struct {
	headNode *ListNode //头节点
}

func main() {

	headNode := &ListNode{}
	headNode1 := &ListNode{}

	listData := headNode
	listData1 := headNode1
	Insert(1, listData, headNode)
	Insert(4, listData, headNode)
	Insert(5, listData, headNode)

	Insert(3, listData1, headNode1)

	cur := make([]*ListNode, 0)
	cur = append(cur, listData)
	cur = append(cur, listData1)

	fmt.Println(cur)
	MergeKLists(cur)
	//fmt.Println(cur)
	//PrintList(listData.Next

}

// 尝试下计算容器的水量
//
//func MaxArray(data []int) {
//
//	var (
//		nowFlg   = sumData(data[1], data[0])
//		rightFlg = sumData(data[1], data[0])
//	)
//
//	for i := 2; i < len(data); i++ {
//
//		rightFlg += data[0]
//	}
//
//}

func sumData(data1, data2 int) (sum int) {
	sum = data1 - data2

	if sum > 0 {
		return sum
	}

	return -sum
}

func MergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	sorted := &ListNode{Val: 0, Next: nil}
	arr := []int{}

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			if lists[i].Next != nil {
				arr = append(arr, lists[i].Val)
				arr = ArrList(lists[i].Next, arr)
			} else {
				arr = append(arr, lists[i].Val)
			}
		}
	}

	fmt.Println(arr)

	n := len(arr)
	if n <= 0 {
		return head.Next
	}
	for i := (n - 1) / 2; i >= 0; i-- {
		ShiftDownHeap(i, n, arr)
	}
	// 顶位末位互换
	for i := n - 1; i > 0; i-- {
		SwapArr(0, i, arr)
		n--
		ShiftDownHeap(0, n, arr)
	}
	for i := 0; i < len(arr); i++ {
		Insert(arr[i], sorted, head)

	}
	return head
}

// ShiftDownHeap ...
func ShiftDownHeap(k int, count int, data []int) {
	for {
		if 2*k+1 >= count {
			break
		}
		j := 2*k + 1
		if j+1 < count && data[j+1] > data[j] {
			j += 1
		}
		if data[k] >= data[j] {
			break
		}
		SwapArr(j, k, data)
		k = j
	}
}

// SwapArr ...
func SwapArr(j int, i int, data []int) {
	t := data[j]
	data[j] = data[i]
	data[i] = t
}

// 链表变成数组 ...
func ArrList(list *ListNode, arr []int) []int {
	if list.Next == nil {
		arr = append(arr, list.Val)
		return arr
	} else {
		arr = append(arr, list.Val)
		arr = ArrList(list.Next, arr)
	}
	return arr
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

//删除节点
func Delete(value int, list *ListNode) {
	pPrev := FindPrevious(value, list)
	_ = pPrev
	p := Find(value, list)
	pPrev = p.Next
	p = nil
}

// FindPrevious ...
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
