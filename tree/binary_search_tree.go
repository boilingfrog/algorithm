package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insertNode(root, newNode *Node) {
	if newNode.value < root.value {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.value > root.value {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}

// Insert 插入一个元素.
func Insert(root *Node, value int) {
	newNode := &Node{value, nil, nil}
	if root == nil {
		root = newNode
	} else {
		insertNode(root, newNode)
	}
}

// 查找一个元素
func Find(root *Node, value int) bool {
	if root == nil {
		return false
	}
	search := root
	for search != nil {
		if search.value == value {
			return true
		}

		if value > search.value {
			search = search.right
		} else {
			search = search.left
		}
	}

	return false
}

// Delete 删除一个元素.
func Delete(root *Node, value int) bool {
	_, existed := remove(root, value)
	return existed
}

// 用来递归移除节点的辅助方法.
// 返回替换root的新节点，以及元素是否存在
func remove(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	// 从左边找
	if value < root.value {
		root.left, existed = remove(root.left, value)
		return root, existed
	}
	// 从右边找
	if value > root.value {
		root.right, existed = remove(root.right, value)
		return root, existed
	}
	// 如果此节点正是要移除的节点,那么返回此节点，同时返回之前可能需要调整.
	existed = true
	// 如果此节点没有孩子，直接返回即可
	if root.left == nil && root.right == nil {
		root = nil
		return root, existed
	}
	// 如果左子节点为空, 提升右子节点
	if root.left == nil {
		root = root.right
		return root, existed
	}
	// 如果右子节点为空, 提升左子节点
	if root.right == nil {
		root = root.left
		return root, existed
	}
	// 如果左右节点都存在,那么从右边节点找到一个最小的节点提升，这个节点肯定比左子树所有节点都大.
	// 也可以从左子树节点中找一个最大的提升，道理一样.
	smallestInRight, _ := Min(root.right)
	// 提升
	root.value = smallestInRight
	// 从右边子树中移除此节点
	root.right, _ = remove(root.right, smallestInRight)
	return root, existed
}

// 最小的值
func Min(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}
	min := root.value
	for root != nil {
		min = root.value
		root = root.left
	}

	return min, true
}

// 最大的值
func Max(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}
	max := root.value
	for root != nil {
		max = root.value
		root = root.right
	}

	return max, true
}

func main() {
	// 头结点
	bst := &Node{
		value: 13,
		left:  nil,
		right: nil,
	}
	Insert(bst, 6)
	Insert(bst, 8)
	Insert(bst, 2)
	Insert(bst, 1)
	Insert(bst, 88)
	Insert(bst, 12)
	Insert(bst, 14)
	fmt.Println(bst.value)
	fmt.Println("-----------")
	fmt.Print("查找值是是否存在：")
	fmt.Println(Find(bst, 12))
	fmt.Println("-----------")
	fmt.Print("最小值：")
	fmt.Println(Min(bst))
	fmt.Println("-----------")
	fmt.Print("最大值：")
	fmt.Println(Max(bst))
	fmt.Println("-----------")
}
