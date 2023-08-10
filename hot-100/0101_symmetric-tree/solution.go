/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 14:21:24
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 14:43:27
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
