package leetcode

/*
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


进阶：

你可以运用递归和迭代两种方法解决这个问题吗？


地址：https://leetcode-cn.com/problems/symmetric-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return compaire(root.Left, root.Right)
}

func compaire(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) && compaire(left.Left, right.Right) && compaire(left.Right, right.Left)
}
