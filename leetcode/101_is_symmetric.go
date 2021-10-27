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

// 使用迭代去做
func isSymmetricD(root *TreeNode) bool {
	if root == nil {
		return false
	}
	leftQ := []*TreeNode{root.Left}
	rightQ := []*TreeNode{root.Right}

	for len(leftQ) > 0 || len(rightQ) > 0 {
		if len(leftQ) != len(rightQ) {
			return false
		}
		left := leftQ[len(leftQ)-1]
		right := rightQ[len(rightQ)-1]

		leftQ = leftQ[:len(leftQ)-1]
		rightQ = rightQ[:len(rightQ)-1]

		if right == nil && left == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		leftQ = append(leftQ, left.Left)
		leftQ = append(leftQ, left.Right)

		rightQ = append(rightQ, right.Right)
		rightQ = append(rightQ, right.Left)
	}
	return true
}
