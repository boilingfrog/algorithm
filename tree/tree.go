package main

type Object interface{}

type TreeNode struct {
	Data       Object
	LeftChild  *TreeNode
	RightChild *TreeNode
}

//(完全)二叉树结构
type Tree struct {
	RootNode *TreeNode
}

//追加元素 (广度优先，即按层级遍历后添加)
func (this *Tree) Add(object Object) {
	node := &TreeNode{Data: object}
	if this.RootNode == nil {
		this.RootNode = node
		return
	}
	queue := []*TreeNode{this.RootNode}
	for len(queue) != 0 {
		cur_node := queue[0]
		queue = queue[1:]

		if cur_node.LeftChild == nil {
			cur_node.LeftChild = node
			return
		} else {
			queue = append(queue, cur_node.LeftChild)
		}
		if cur_node.RightChild == nil {
			cur_node.RightChild = node
			return
		} else {
			queue = append(queue, cur_node.RightChild)
		}
	}
}

//广度遍历
func (this *Tree) BreadthTravel() {

	if this.RootNode == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, this.RootNode)

	for len(queue) != 0 {
		//fmt.Printf("len(queue):%d\n", len(queue))
		cur_node := queue[0]
		queue = queue[1:]

		if cur_node.LeftChild != nil {
			queue = append(queue, cur_node.LeftChild)
		}
		if cur_node.RightChild != nil {
			queue = append(queue, cur_node.RightChild)
		}
	}

}
