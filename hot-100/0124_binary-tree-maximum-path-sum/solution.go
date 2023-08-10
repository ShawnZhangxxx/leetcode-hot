/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 10:10:29
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-03 11:04:16
 * @Description:
 */

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	dfs(root, &maxSum)
	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, maxSum)
	r := dfs(root.Right, maxSum)
	*maxSum = max(*maxSum, root.Val+l+r, root.Val+l, root.Val+r, root.Val)
	return max(root.Val, root.Val+l, root.Val+r)
}

func max(nums ...int) int {
	res := math.MinInt64
	for _, item := range nums {
		if item > res {
			res = item
		}
	}

	return res
}
