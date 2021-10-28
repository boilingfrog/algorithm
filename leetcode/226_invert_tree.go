package leetcode

/*

226. 翻转二叉树
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

地址：https://leetcode-cn.com/problems/invert-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 这是官方给出的答案
// 害，自己开始做有点误入歧途，想着把 树中根节点开始计算分成左右两部分，然后递归的更换左子树的左节点和右子树的又节点
// 当然这种也不是不行，仔细想了想确实，只有在左右完美对称的情况下，才能完成，否则只有一个地方为nil，就不能成功替换了

// 看了官方的答案
// 首先更换左右子树，然后在左子树和右子树分别对自己的左右节点进行左右替换
// 这种确实是妙
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
