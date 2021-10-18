package leetcode

/*
给你二叉树的根节点 root ，返回它节点值的前序遍历。

示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]


示例 2：

输入：root = []
输出：[]


示例 3：

输入：root = [1]
输出：[1]



示例 4：

输入：root = [1,2]
输出：[1,2]


示例 5：

输入：root = [1,null,2]
输出：[1,2]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

地址：https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 首先我们知道前序便利的先后顺序是：先打印这个节点，然后再打印它的左子树，最后打印它的右子树

// 使用递归
// 总结：递归还是比较难写的，先写出一个运算，然后找到对应的关系，找到退出的条件

// 错误的思考：使用递归，纠结了好久，每次总想一下子，把所有的递归流程全部想出来，但是想的越是多就越写不出来，容易进入死胡同

// 前序的就是先当前节点，然后左节点，然后右节点，这就是一层的递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}

// 迭代法
// 迭代法引入了栈
// 先入栈，先右子树，然后是左子树，因为栈是先进后出的
// 每次从栈定去取出node，接者处理，因为左节点最后进入，所以每次总是最先处理左节点
func preorderTraversalIteration(root *TreeNode) []Object {
	var res []Object
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}
