package leetcode

import "fmt"

/*
257. 二叉树的所有路径
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。


示例 1：

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]


示例 2：

输入：root = [1]
输出：["1"]


提示：

树中节点的数目在范围 [1, 100] 内
-100 <= Node.val <= 100


地址：https://leetcode-cn.com/problems/binary-tree-paths/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 总结：这道题倒是自己做出来了
// 时间复杂度0(N的平方)
// 空间复杂度0(N的平方)
var sil []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	sil = []string{}
	if root.Left != nil {
		binaryTreePathsRec(root.Left, fmt.Sprintf("%d->%d", root.Val, root.Left.Val))
	}

	if root.Right != nil {
		binaryTreePathsRec(root.Right, fmt.Sprintf("%d->%d", root.Val, root.Right.Val))
	}

	return sil
}

func binaryTreePathsRec(root *TreeNode, s string) {
	if root.Left == nil && root.Right == nil {
		sil = append(sil, s)
		return
	}
	if root.Left != nil {
		s1 := s + fmt.Sprintf("->%d", root.Left.Val)
		binaryTreePathsRec(root.Left, s1)
	}
	if root.Right != nil {
		s1 := s + fmt.Sprintf("->%d", root.Right.Val)
		binaryTreePathsRec(root.Right, s1)
	}

	return
}
