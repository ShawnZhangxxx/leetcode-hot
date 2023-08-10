/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 14:00:44
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 14:38:43
 * @Description:
 */

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if left >= root.Val || root.Val >= right {
		return false
	}

	return validate(root.Left, left, root.Val) && validate(root.Right, root.Val, right)
}
