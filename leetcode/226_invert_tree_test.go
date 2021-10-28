package leetcode

import "testing"

func TestInvertTree(t *testing.T) {
	node := &TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val: 4,
		},
	}
	t.Log(node.Val, node.Left, node.Right)

	node1 := invertTree(node)

	t.Log(node1.Val, node1.Left, node1.Right)

}
