package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

// BST 是一个节点的值为int类型的二叉搜索树.
type BST struct {
	root *Node
}

// 从根节点依次比较
func insertNode(root, newNode *Node) {
	if newNode.value < root.value { // 应该放到根节点的左边
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.value > root.value { // 应该放到根节点的右边
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
	// 否则等于根节点
}

// Insert 插入一个元素.
func (bst *BST) Insert(value int) {
	newNode := &Node{value, nil, nil}
	// 如果二叉树为空，那么这个节点就当作跟节点
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

// 查找一个元素
func (bst *BST) Find(value int) bool {
	if bst.root == nil {
		return false
	}
	search := bst.root
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

func main() {
	bst := BST{}
	bst.Insert(13)
	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(88)
	bst.Insert(12)
	bst.Insert(14)
	fmt.Println("-----------")
	fmt.Println(bst.Find(12))
	fmt.Println("-----------")
}
