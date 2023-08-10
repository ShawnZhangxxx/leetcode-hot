/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 15:45:19
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 15:58:11
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 迭代
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	q := []*TreeNode{root}
	for len(q) != 0 {
		size := len(q)
		level++
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return level
}
