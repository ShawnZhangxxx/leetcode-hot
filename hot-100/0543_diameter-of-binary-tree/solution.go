/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 09:58:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 10:07:42
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	depth(root, &diameter)
	return diameter
}

func depth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}

	l := depth(node.Left, diameter)
	r := depth(node.Right, diameter)
	*diameter = max(*diameter, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
