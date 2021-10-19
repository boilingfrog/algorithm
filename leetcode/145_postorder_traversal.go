package leetcode

/*
给定一个二叉树，返回它的 后序遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]

地址：https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 使用递归确实是太爽了
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, postorderTraversal(root.Left)...)
		res = append(res, postorderTraversal(root.Right)...)
		res = append(res, root.Val)
	}

	return res
}
