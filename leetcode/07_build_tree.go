package leetcode

/*

剑指 Offer 07. 重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



示例 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
示例 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]


限制：

0 <= 节点个数 <= 5000


地址：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 答题分析

// 这道题目自己没有相处答案，还会看了官方的题解

//
/*

首先需要你弄明白前序遍历的结果
	二叉树前序遍历的顺序为：
		1、先遍历根节点；
		2、随后递归地遍历左子树；
		3、最后递归地遍历右子树。

	二叉树中序遍历的顺序为：
		1、先递归地遍历左子树；
		2、随后遍历根节点；
		3、最后递归地遍历右子树。

通过分析前序和后续遍历的结果可以知道

1、前序的第一个节点肯定是顶节点

2、中序节点遍历的结果中，顶节点的左边为左子树，顶节点右边为右子树

3、思路，找到顶节点，区分开左子树和右子树，左，右子树的前序的第一个节点又是对应子树的根节点

4、上面已经把这个问题划分成不同的子问题，递归的去处理就行了

5、实在是妙


时间复杂度：O(n)

空间复杂度：O(n)
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[0:i+1])], inorder[:i+1])

	root.Right = buildTree(preorder[len(inorder[0:i+1]):], inorder[i+1:])

	return root
}
