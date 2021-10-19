package leetcode

/*
94. 二叉树的中序遍历

给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：

输入：root = [1,null,2,3]
输出：[1,3,2]


示例 2：

输入：root = []
输出：[]


示例 3：

输入：root = [1]
输出：[1]


示例 4：

输入：root = [1,2]
输出：[2,1]


示例 5：

输入：root = [1,null,2]
输出：[1,2]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


地址：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 使用递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

// 使用迭代
// 这次的迭代又没有写出来，好气
// 梳理下原因 因为迭代使用的栈进行处理，如何巧妙的出入栈这块刚开始没想明白，看了答案豁然开朗，实在是妙
func inorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}

	return res
}
