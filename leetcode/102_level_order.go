package leetcode

/*
102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

地址：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 使用迭代来处理，谁虽然迭代是相对简单的，不过也写了好久，有点尴尬
// 卡在了对数组的处理上了
var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return [][]int{}
	}
	return levelOrderIter(root, -1)
}

func levelOrderIter(root *TreeNode, num int) [][]int {
	if root != nil {
		num++
		if len(res) < num+1 {
			res = append(res, []int{})
		}
		res[num] = append(res[num], root.Val)
		levelOrderIter(root.Left, num)
		levelOrderIter(root.Right, num)
	}
	return res
}
