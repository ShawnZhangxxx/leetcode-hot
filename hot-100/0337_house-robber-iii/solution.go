/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-25 10:25:47
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-25 11:04:15
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(robHelper(root))
}

func robHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	money := root.Val
	l1, l2 := robHelper(root.Left)
	r1, r2 := robHelper(root.Right)
	// 抢劫当前节点的最大金额
	c1 := money + l2 + r2
	// 不抢劫当前节点的最大金额
	c2 := max(l1, l2) + max(r1, r2)
	return c1, c2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
